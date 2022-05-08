package parser

import "fmt"

//Parse @apiHeaderExample Content
func Api_headerExample(content string, log ...bool) *ApiExampleStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiHeaderExample with Content: ", content)
	}

	result := Api_example(content)

	return result
}
