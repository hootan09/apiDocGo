package parser

import (
	"fmt"
	"strings"
)

// Parse @apiVersion Content
func Api_version(content string, log ...bool) string {
	if len(log) > 0 {
		fmt.Println("Parse @apiVersion with Content: ", content)
	}
	result := strings.TrimSpace(content)
	return result
}
