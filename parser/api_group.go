package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type ApiGroupStruct struct {
	Title string `json:"title"`
}

// Parse @apiGroup content
// Return Struct with parameters {Title}
func Api_group(content string, log ...bool) *ApiGroupStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiGroup with Content: ", content)
	}

	trimContent := strings.TrimSpace(content)

	r := regexp.MustCompile(`(\s+)`)

	match := r.ReplaceAllString(trimContent, "_")

	apiGroup := new(ApiGroupStruct)

	apiGroup.Title = match

	return apiGroup
}
