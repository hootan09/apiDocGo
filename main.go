package main

import (
	"ApiDocGo/parser"
	"fmt"
)

func main() {
	Documents := parser.ReadDoc("./examples")
	fmt.Printf("%+v \n", Documents)

	// result := parser.Api("{get} /user/:region/:id/:opt Read data of a User", true)
	// fmt.Printf("%+v \n", result)

	// json_data, err := json.Marshal(Documents)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(json_data))

	//TODO
	// usage "apidocgo -i routes/ -o public/apidoc"

}
