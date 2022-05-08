package parser

import "fmt"

//Parse @apiHeader Content
func Api_header(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiHeader with Content: ", content)
	}

	result := Api_param(content)
	result.Group = "Header"
	return result
}
