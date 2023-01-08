package api

import (
	"context"
	"fmt"
	"github.com/grizmar-realty/scrape_proxy/internal/http_client"
	"github.com/grizmar-realty/scrape_proxy/internal/parsers"
)

const CaptchaExpr = "(<form.method=.*?action=.*?id=.form_captcha.>)"

var CaptchaError = fmt.Errorf("captcha error")

type Implementation struct {
	client *http_client.Client

	UnimplementedScrapeProxyServer
}

func NewImplementation(client *http_client.Client) *Implementation {
	return &Implementation{
		client: client,
	}
}

func (i *Implementation) SendRequest(ctx context.Context, r *Request) (*Response, error) {
	request, err := http_client.NewHttpRequest(r.Url, r.HttpMethod)
	if err != nil {
		return nil, fmt.Errorf("SendRequest: NewHttpRequest: %w", err)
	}

	resp, err := i.client.SendRequest(request)
	if err != nil {
		return nil, fmt.Errorf("SendRequest: client.SendRequest: %w", err)
	}

	if err := i.CheckCaptcha(string(resp.Body)); err != nil {
		return nil, err
	}

	body, err := i.ParseResponse(ctx, string(resp.Body), r.ParsePatterns)

	return &Response{
		StatusCode: int32(resp.Status),
		Body:       []byte(body),
	}, nil
}

func (i *Implementation) CheckCaptcha(data string) error {
	captchaParser, err := parsers.NewHtmlParser(CaptchaExpr)
	if err != nil {
		return err
	}
	_, err = captchaParser.Parse(data)
	if err == nil {
		return CaptchaError
	}

	return nil
}

func (i *Implementation) ParseResponse(ctx context.Context, body string, parsePatterns []*ParseInfo) (string, error) {
	result := body

	for _, parseInfo := range parsePatterns {
		var (
			parser parsers.Parser
			err    error
		)

		switch parseInfo.Type {
		case ParserType_HTML:
			parser, err = parsers.NewHtmlParser(parseInfo.GetExpr())
			if err != nil {
				return "", fmt.Errorf("NewHtmlParser: %w", err)
			}
		case ParserType_JSON:
			parser = parsers.NewJsonParser(parseInfo.GetExpr())
		default:
			return "", fmt.Errorf("unknown parser type")
		}

		result, err = parser.Parse(result)
		if err != nil {
			return "", fmt.Errorf("parserType=%v: parse error: %w", parseInfo.Type, err)
		}
	}

	return result, nil
}
