package parser

import "fmt"

//Parse @apiSuccessExample Content
func Api_successExample(content string, log ...bool) *ApiExampleStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiSuccessExample with Content: ", content)
	}

	result := Api_example(content)
	return result
}
