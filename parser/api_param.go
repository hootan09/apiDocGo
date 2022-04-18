package parser

import (
	"fmt"
)

type ApiParamStruct struct {
	// Title string `json:"title"`
}

// Parse @apiName content
// Return Struct with parameters {}
func Api_param(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiParam with Content: ", content)
	}

	// trimContent := strings.TrimSpace(content)

	// r := regexp.MustCompile(`(\s+)`)

	// match := r.ReplaceAllString(trimContent, "_")

	apiParam := new(ApiParamStruct)

	// apiParam.Title = match

	return apiParam
}
