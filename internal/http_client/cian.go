package http_client

import (
	"compress/gzip"
	"fmt"
	http "github.com/useflyent/fhttp"
	"io"
	"log"
)

// Client TODO сделать тесты на клиент
type Client struct {
	httpClient http.Client
}

type Response struct {
	Body   []byte
	Status int
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{
			Transport: &http.Transport{},
		},
	}
}

func NewHttpRequest(url string, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewHttpRequest: %w", err)
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-encoding", "deflate, br")
	req.Header.Add("Accept-language", "ru-RU,ru;q=0.5")
	//TODO добавить разные юзер агенты
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:103.0) Gecko/20100101 Firefox/103.0")

	return req, nil
}

func (c *Client) SendRequest(req *http.Request) (*Response, error) {
	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error: %w", err)
	}

	result := &Response{
		Status: httpResponse.StatusCode,
	}

	if httpResponse.StatusCode == 200 {
		var reader io.ReadCloser
		switch httpResponse.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(httpResponse.Body)
			if err != nil {
				log.Println(err)

				//TODO разобраться почему не получается распарсить gzip если приходит captcha
				reader = httpResponse.Body
			}

			defer reader.Close()
		default:
			reader = httpResponse.Body
		}

		body, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("read response body error: %w", err)
		}

		result.Body = body
	}

	return result, nil
}
