package main

import (
	"ApiDocGo/cmd"
	"ApiDocGo/parser"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	ApiDocGoVersion = "0.0.1"
)

func main() {
	/* parse flags */

	port := cmd.ParseFlags()
	Documents := parser.ReadDoc("./examples") //must be dynamic path by user
	// json_data, err := json.Marshal(*Documents)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(json_data))

	parser.WriteDoc(*Documents, ApiDocGoVersion)

	log.Println("Build Done!")

	// if serve static file is true
	if port > 0 {
		wdPath, _ := os.Getwd()
		buildPath := path.Join(wdPath, parser.BuildFolderName)
		fs := http.FileServer(http.Dir(buildPath))
		http.Handle("/", fs)
		log.Printf("Listening on http://localhost:%d", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
