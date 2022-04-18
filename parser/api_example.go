package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type ApiExampleStruct struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Parse @apiExample content
// Return Struct with parameters {Title, Type, Content}
func Api_example(content string, log ...bool) *ApiExampleStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiExample with Content: ", content)
	}

	trimContent := strings.TrimSpace(content)

	r := regexp.MustCompile(`(@\w*)?(?:(?:\s*\{\s*([a-zA-Z0-9./\\[\]_-]+)\s*\}\s*)?\s*(.*)?)?`)

	match := r.FindAllStringSubmatch(trimContent, -1)

	apiExample := new(ApiExampleStruct)

	apiExample.Type = match[0][2]
	apiExample.Title = match[0][3]

	matchContent := r.FindAllString(trimContent, -1)

	apiExample.Content = strings.TrimSpace(strings.Join(matchContent[1:], "\n"))

	return apiExample
}
