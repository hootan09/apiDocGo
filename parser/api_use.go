package parser

import (
	"fmt"
	"strings"
)

//Parse @apiUse Content
func Api_use(content string, log ...bool) string {
	if len(log) > 0 {
		fmt.Println("Parse @apiUse with Content: ", content)
	}

	result := strings.TrimSpace(content)
	return result
}
