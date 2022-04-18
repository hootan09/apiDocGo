package parser

import "fmt"

//Parse @apiBody Content
func Api_body(content string, log ...bool) {
	if len(log) > 0 {
		fmt.Println("Parse @apiBody with Content: ", content)
	}

	//Todo
	//Needed @apiParams First
}
