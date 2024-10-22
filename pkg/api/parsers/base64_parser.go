package parsers

import (
	"encoding/base64"
	"fmt"
)

type Base64Parser struct{}

func NewBase64Parser() *Base64Parser {
	return &Base64Parser{}
}

func (p *Base64Parser) Parse(data string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}

	return string(res), nil
}
