package parsers

import (
	"fmt"
	"regexp"
)

type HtmlParser struct {
	reg *regexp.Regexp
}

func NewHtmlParser(expr string) (*HtmlParser, error) {
	reg, err := regexp.Compile(expr)
	if err != nil {
		return nil, fmt.Errorf("regexp compile error: %w", err)
	}

	return &HtmlParser{
		reg: reg,
	}, nil
}

func (p *HtmlParser) Parse(data string) (string, error) {
	result := p.reg.FindStringSubmatch(data)

	if len(result) < 2 {
		return "", fmt.Errorf("not found data from html string")
	}

	return result[1], nil
}
