package parser

import "fmt"

//Parse @apiSuccess Content
func Api_success(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiSuccess with Content: ", content)
	}

	result := Api_param(content)
	result.Group = "Success 200"
	return result
}
