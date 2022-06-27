package main

import (
	"ApiDocGo/cmd"
	"ApiDocGo/parser"
)

const (
	ApiDocGoVersion = "0.0.1"
)

func main() {
	//parse flags
	cmd.ParseFlags() //hasServe := cmd.ParseFlags() if (hasserve == true) must run http static server with cmd.Port

	Documents := parser.ReadDoc("./examples") //must be dynamic path by user

	// json_data, err := json.Marshal(*Documents)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(json_data))

	parser.WriteDoc(*Documents, ApiDocGoVersion)

}
