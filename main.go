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

	// content := `GetCategory`
	// result := parser.Api_name(content)

	// content := `{Number} id <code>id</code> of the user.`
	// content := `{String{1..4}} [user.name='John Doe'] Users fullname.`
	// result := parser.Api_param(content)

	// content := `{String} name Name of the User.`
	// result := parser.Api_body(content)

	// content := ` Response (example):
	// *     HTTP/1.1 401 Not Authenticated
	// *     {
	// *       "error": "NoAccessRight"
	// *     }`
	// result := parser.Api_errorExample(content)

	// content := `(500 Internal Server Error) InternalServerError The server encountered an internal error.`
	// result := parser.Api_error(content)

	// content := ` {Header} Header-Example
	// *     "Authorization: token 5f048fe"`
	// result := parser.Api_headerExample(content)

	// content := `{String} Authorization The token can be generated from your user profile.`
	// result := parser.Api_header(content)

	// content := `{json} Some json code:
	// *   {
	// *     "user": "Sample User",
	// *      "payload": {
	// *        "test": [
	// *          "code": "
	// *            public class HelloWorldTest {
	// *              HelloWorld hw = new HelloWorld();
	// *              @Test
	// *              public void testOkay {
	// *              assertEquals(\"HelloWorld\", hw.getMsg());
	// *            }
	// *         }"
	// *        ]
	// *      }
	// *   }`
	// result := parser.Api_paramExample(content)

	// content := `{String=Aerial,Land,Underwater} view=Aerial Type of view.`
	// result := parser.Api_query(content)

	content := `{json} Success-Example
	*     HTTP/1.1 200 OK
	*     {
	*         "result": "ok"
	*     }`
	result := parser.Api_successExample(content)
	fmt.Printf("%+v\n", result)

}
