package main

import (
	"ApiDocGo/parser"
	"encoding/json"
	"fmt"
)

const (
	ApiDocGoVersion = "0.0.1"
)

func main() {
	//TODO
	// usage "apidocgo -i routes/ -o public/apidoc"

	Documents := parser.ReadDoc("./examples")

	json_data, err := json.Marshal(Documents)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json_data))

	parser.WriteDoc(Documents, ApiDocGoVersion)

}
