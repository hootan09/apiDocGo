package main

import (
	"ApiDocGo/parser"
	"fmt"
)

func main() {
	Documents := parser.ReadDoc("./examples")
	fmt.Printf("%+v \n", Documents)

	// json_data, err := json.Marshal(Documents)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(json_data))

	//TODO
	// usage "apidocgo -i routes/ -o public/apidoc"

}
