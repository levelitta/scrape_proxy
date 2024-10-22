package api

//go:generate mockery --all --with-expecter --case snake

import (
	"context"
	"fmt"
	"github.com/levelitta/scrape_proxy/pkg/api/http_client"
	"github.com/levelitta/scrape_proxy/pkg/api/parsers"
	"log"
)

const CaptchaExpr = "(<form.method=.*?action=.*?id=.form_captcha.>)"

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
	fn := "Implementation.SendRequest"

	log.Printf("%s: url: %s", fn, r.GetUrl())
	log.Printf("%s: body: %s", fn, r.GetBody())

	request, err := http_client.NewHttpRequest(r.GetUrl(), r.GetHttpMethod(), r.GetHeaders(), r.GetBody())
	if err != nil {
		log.Printf("%s: NewHttpRequest: error=%s", fn, err)
		return nil, fmt.Errorf("%s: NewHttpRequest: %w", fn, err)
	}

	log.Printf("%s: request: %v", fn, request)

	resp, err := i.client.SendRequest(request)
	if err != nil {
		log.Printf("%s: client.SendRequest: error=%s", fn, err)
		return nil, fmt.Errorf("%s: client.SendRequest: %w", fn, err)
	}

	if err := i.CheckCaptcha(string(resp.Body)); err != nil {
		log.Printf("%s: %s", fn, err)
		return nil, err
	}

	body, err := i.ParseResponse(string(resp.Body), r.ParsePatterns)
	if err != nil {
		log.Printf("%s: ParseResponse: error=%s", fn, err)
		return nil, fmt.Errorf("%s: ParseResponse: %w", fn, err)
	}

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
		return NewCaptchaError()
	}

	return nil
}

func (i *Implementation) ParseResponse(body string, parsePatterns []*ParseInfo) (string, error) {
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
		case ParserType_Base64:
			parser = parsers.NewBase64Parser()
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
