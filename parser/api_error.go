package parser

import "fmt"

const (
	error4xx = "Error 4xx"
	error5xx = "Error 5xx"
)

//Parse @apiError Content
func Api_error(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiError with Content: ", content)
	}

	result := Api_param(content)
	result.Group = error4xx
	return result
}
