package parser

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	allowedValuesWithDoubleQuoteRegExp = `/"[^"]*[^"]"`
	allowedValuesWithQuoteRegExp       = "/'[^']*[^']'"
	allowedValuesRegExp                = `/[^,\s]+`
)

type ApiParamStruct struct {
	Group         string   `json:"group"`
	Type          string   `json:"type"`
	Size          string   `json:"size"`
	AllowedValues []string `json:"allowedValues"`
	Optional      bool     `json:"optional"`
	Field         string   `json:"field"`
	DefaultValue  string   `json:"defaultVale"`
	Description   string   `json:"description"`
}

// Search: group, type, optional, fieldname, defaultValue, size, description
// Example: {String{1..4}} [user.name='John Doe'] Users fullname.
//
// Naming convention:
//     b -> begin
//     e -> end
//     name -> the field value
//     oName -> wrapper for optional field
//     wName -> wrapper for field
var regExp = map[string]interface{}{
	"oGroup": map[string]interface{}{
		"b":     `\s*(?:\(\s*`, // starting with '(', optional surrounding spaces
		"group": `(.+?)`,       // 1
		"e":     `\s*\)\s*)?`,  // ending with ')', optional surrounding spaces
	},
	"oType": map[string]interface{}{ // optional type: {string}
		"b":    `\s*(?:\{\s*`,                     // starting with '{', optional surrounding spaces
		"type": `([a-zA-Z0-9()#:\.\/\\\[\]_|-]+)`, // 2
		"oSize": map[string]interface{}{ // optional size within type: {string{1..4}}
			"b":    `\s*(?:\{\s*`, // starting with '{', optional surrounding spaces
			"size": `(.+?)`,       // 3
			"e":    `\s*\}\s*)?`,  // ending with '}', optional surrounding spaces
		},
		"oAllowedValues": map[string]interface{}{ // optional allowed values within type: {string='abc','def'}
			"b":              `\s*(?:=\s*`,      // starting with '=', optional surrounding spaces
			"possibleValues": `(.+?)`,           // 4
			"e":              `(?:=\s*\}\s*))?`, // ending with '}', optional surrounding spaces
		},
		"e": `\s*\}\s*)?`, // ending with '}', optional surrounding spaces
	},
	"wName": map[string]interface{}{
		"b":         `(\[?\s*`,                        // 5 optional optional-marker
		"name":      `([#@a-zA-Z0-9\$\:\.\/\\_-]+`,    // 6
		"withArray": `(?:\[[a-zA-Z0-9\.\/\\_-]*\])?)`, // https://github.com/apidoc/apidoc-core/pull/4
		"oDefaultValue": map[string]interface{}{ // optional defaultValue
			"b":               `(?:\s*=\s*(?:`,     // starting with '=', optional surrounding spaces
			"withDoubleQuote": `"([^"]*)"`,         // 7
			"withQuote":       `|'([^']*)'`,        // 8
			"withoutQuote":    `|(.*?)(?:\s|\]|$)`, // 9
			"e":               `))?`,
		},
		"e": `\s*\]?\s*)`,
	},
	"description": `(.*)?`, // 10
}

// Parse @apiName content
// Return Struct with parameters {Group,Type,Size,AllowedValues,Optional,Field,DefaultValue,Description}
func Api_param(content string, log ...bool) *ApiParamStruct {
	if len(log) > 0 {
		fmt.Println("Parse @apiParam with Content: ", content)
	}

	trimContent := strings.TrimSpace(content)

	// replace Linebreak with Unicode
	trimContent = strings.ReplaceAll(trimContent, "\n", "\uffff")

	// regularExpressions := IterMap(regExp)
	// r := regexp.MustCompile(regularExpressions)

	//TODO:
	// myCode Do not give back same regex (must be fixed)
	regularExpressions := `\s*(?:\(\s*(.+?)\s*\)\s*)?\s*(?:\{\s*([a-zA-Z0-9()#:\.\/\\\[\]_|-]+)\s*(?:\{\s*(.+?)\s*\}\s*)?\s*(?:=\s*(.+?)(?:=\s*\}\s*))?\s*\}\s*)?(\[?\s*([#@a-zA-Z0-9\$\:\.\/\\_-]+(?:\[[a-zA-Z0-9\.\/\\_-]*\])?)(?:\s*=\s*(?:"([^"]*)"|'([^']*)'|(.*?)(?:\s|\]|$)))?\s*\]?\s*)(.*)?`
	r := regexp.MustCompile(regularExpressions)

	match := r.FindAllStringSubmatch(trimContent, -1)

	// reverse Unicode Linebreaks
	for i, v := range match {
		for j, _ := range v {
			strings.ReplaceAll(match[i][j], "\uffff", "\n")
		}
	}

	apiParam := new(ApiParamStruct)
	allowedValues := match[0][4]
	if allowedValues != "" {
		if string(allowedValues[0]) == `"` {
			regularExpressions = allowedValuesWithDoubleQuoteRegExp
		} else if string(allowedValues[0]) == "'" {
			regularExpressions = allowedValuesWithQuoteRegExp
		} else {
			regularExpressions = allowedValuesRegExp
		}
	}
	m := regexp.MustCompile(regularExpressions)
	allowedValuesMatch := m.FindAllStringSubmatch(allowedValues, -1)

	if len(allowedValuesMatch) > 0 {
		apiParam.AllowedValues = allowedValuesMatch[0] //Must be Checked
	}

	if apiParam.Group = match[0][1]; apiParam.Group == "" {
		apiParam.Group = "Parameter"
	}

	apiParam.Type = match[0][2]
	apiParam.Size = match[0][3]
	apiParam.Optional = len(match[0][5]) > 0 && string(match[0][5][0]) == "["
	apiParam.Field = match[0][6]

	apiParam.DefaultValue = match[0][7] + match[0][8] + match[0][9]

	apiParam.Description = match[0][10]

	// fmt.Println(`0`, match[0][0])
	// fmt.Println(`1`, match[0][1])
	// fmt.Println(`2`, match[0][2])
	// fmt.Println(`3`, match[0][3])
	// fmt.Println(`4`, match[0][4])
	// fmt.Println(`5`, match[0][5])
	// fmt.Println(`6`, match[0][6])
	// fmt.Println(`7`, match[0][7])
	// fmt.Println(`8`, match[0][8])
	// fmt.Println(`9`, match[0][9])
	// fmt.Println(`10`, match[0][10])

	return apiParam
}

// Iterate Nested map with Type of map[string]interface{}
// return value string
func IterMap(m map[string]interface{}) string {
	str := ""
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			str += IterMap(v.(map[string]interface{}))
			continue
		case string:
			str += m[k].(string)
			continue
		default:
			return ""
		}
	}
	return str
}
