package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type ApiDefineStruct struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Parse @apiDefine content
func Api_define(content string, log ...bool) *ApiDefineStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiDefine with Content: ", content)
	}

	trimContent := strings.TrimSpace(content)

	apiDefine := new(ApiDefineStruct)

	r := regexp.MustCompile(`(?m)^([\w:]*)(.*?)(?:\s+|$)(.*)$`)

	match := r.FindAllStringSubmatch(trimContent, -1)

	apiDefine.Name = match[0][1]
	apiDefine.Title = match[0][3]

	matchDescription := r.FindAllString(trimContent, -1)
	description := matchDescription[1:]

	apiDefine.Description = strings.TrimSpace(strings.Join(description, "\n"))

	return apiDefine
}
