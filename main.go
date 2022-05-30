package main

import (
	"ApiDocGo/parser"
)

const (
	ApiDocGoVersion = "0.0.1"
)

func main() {
	//TODO
	// usage "apidocgo -i routes/ -o public/apidoc"

	// json_data, err := json.Marshal(Documents)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(json_data))

	Documents := parser.ReadDoc("./examples")
	parser.WriteDoc(Documents, ApiDocGoVersion)

}
