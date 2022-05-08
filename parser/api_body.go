package parser

import "fmt"

//Parse @apiBody Content
func Api_body(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiBody with Content: ", content)
	}

	result := Api_param(content)
	result.Group = "Body"

	return result
}
