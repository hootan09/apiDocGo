package main

import (
	"ApiDocGo/cmd"
)

const (
	ApiDocGoVersion = "0.0.1"
)

func main() {
	// Documents := parser.ReadDoc("./examples") //must be dynamic path by user

	// json_data, err := json.Marshal(*Documents)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(json_data))

	// parser.WriteDoc(*Documents, ApiDocGoVersion)

	cmd.ParseFlags()

}
