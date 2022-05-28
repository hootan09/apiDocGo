// Package parser provide parsing functions
package parser

import (
	"fmt"
	"regexp"
)

type ApiStruct struct {
	Type  string `json:"type"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

// Parse @api Content
func Api(content string, log ...bool) *ApiStruct {
	if len(log) > 0 {
		fmt.Println("Parse @api with Content: ", content)
	}
	api := new(ApiStruct)

	r := regexp.MustCompile(`^(?:(?:\{(.+?)\})?\s*)?(.+?)(?:\s+(.+?))?$`)

	match := r.FindAllStringSubmatch(content, -1)

	api.Type = match[0][1]
	api.Url = match[0][2]
	api.Title = match[0][3]

	return api
}
