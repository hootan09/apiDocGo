package parser

import "fmt"

//Parse @apiPermission Content
func Api_permission(content string, log ...bool) {
	if len(log) > 0 {
		fmt.Println("Parse @apiPermission with Content: ", content)
	}

	//todo
	//needed @apiUse
}
