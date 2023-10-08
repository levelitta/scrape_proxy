package parsers

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type JsonParser struct {
	expr string
}

func NewJsonParser(expr string) *JsonParser {
	return &JsonParser{expr: expr}
}

func (p *JsonParser) Parse(data string) (string, error) {
	val := gjson.Get(data, p.expr)
	if val.String() == "" {
		return "", fmt.Errorf("results not found")
	}

	return val.String(), nil
}
