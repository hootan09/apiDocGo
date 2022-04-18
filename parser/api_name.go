package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type ApiNameStruct struct {
	Title string `json:"title"`
}

// Parse @apiName content
// Return Struct with parameters {Title}
func Api_name(content string, log ...bool) *ApiNameStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiName with Content: ", content)
	}

	trimContent := strings.TrimSpace(content)

	r := regexp.MustCompile(`(\s+)`)

	match := r.ReplaceAllString(trimContent, "_")

	apiName := new(ApiNameStruct)

	apiName.Title = match

	return apiName
}
