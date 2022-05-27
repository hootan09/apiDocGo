package main

import (
	"ApiDocGo/parser"
	"fmt"
)

func main() {
	Documents := parser.ReadDoc("./examples")
	fmt.Printf("%+v \n", Documents)

	//TODO
	// usage "apidocgo -i routes/ -o public/apidoc"

}
