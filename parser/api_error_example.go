package parser

import "fmt"

//Parse @apiErrorExample Content
func Api_errorExample(content string, log ...bool) *ApiExampleStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiErrorExample with Content: ", content)
	}

	result := Api_example(content)
	return result
}
