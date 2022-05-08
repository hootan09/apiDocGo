package parser

import "fmt"

//Parse @apiParamExample Content
func Api_paramExample(content string, log ...bool) *ApiExampleStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiParamExample with Content: ", content)
	}

	result := Api_example(content)
	return result
}
