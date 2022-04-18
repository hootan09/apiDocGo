package main

import (
	"ApiDocGo/parser"
	"fmt"
)

func main() {
	// result := parser.Api("{post} /api/v1/create Create new article", true)

	// content := `admin:computer User access only
	// * This optional description belong to to the group admin.
	// * This optional description belong to to the new admin.`
	// result := parser.Api_define(content)

	// content := `{bash} Curl example
	// * curl -H "Authorization: token 5f048fe" -i https://api.example.com/user/fr-par/4711
	// * curl -H "Authorization: token 5f048fe" -H "X-Apidoc-Cool-Factor: superbig" -i https://api.example.com/user/de-ber/1337/yep`
	// result := parser.Api_example(content)

	// content := `Category (official)`
	// result := parser.Api_group(content)

	content := `GetCategory`
	result := parser.Api_name(content)
	fmt.Printf("%+v\n", result)
}
