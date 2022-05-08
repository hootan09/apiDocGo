package parser

import "fmt"

//Parse @apiQuery Content
func Api_query(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiQuery with Content: ", content)
	}

	result := Api_param(content)
	result.Group = "Query"
	return result
}
