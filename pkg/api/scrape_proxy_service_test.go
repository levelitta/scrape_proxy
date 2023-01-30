package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementation_ParseResponse(t *testing.T) {
	var testData = `
		<html>
		window._cianConfig['header-frontend'] = (window._cianConfig['header-frontend'] || []).concat([{"result":"offers"}]);
		</html>
	`

	i := Implementation{}

	parsePatterns := []*ParseInfo{
		{
			Type: ParserType_HTML,
			Expr: "concat\\((.*\\])\\);",
		},
		{
			Type: ParserType_JSON,
			Expr: "@this.0.@values",
		},
	}

	res, err := i.ParseResponse(testData, parsePatterns)
	assert.NoError(t, err)

	assert.Equal(t, "[\"offers\"]", res)
}
func TestImplementation_CheckCaptcha(t *testing.T) {
	var testData = `<form method=post action="" id="form_captcha">`

	i := Implementation{}

	err := i.CheckCaptcha(testData)
	assert.Error(t, NewCaptchaError(), err)
}
