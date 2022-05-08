package parser

import (
	"fmt"
	"strings"
)

//Parse @apiSampleRequest Content
func Api_sampleRequest(content string, log ...bool) string {
	if len(log) > 0 {
		fmt.Println("Parse @apiSampleRequest with Content: ", content)
	}

	result := strings.TrimSpace(content)
	return result
}
