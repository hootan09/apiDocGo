package parser

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
)

// if JSON config file not exist must using this
var DefaultConfig = map[string]string{
	"name":        "Acme project",
	"version":     "0.0.0",
	"description": "REST Api Doc",
}

//TODO
//get apidocgo json config file if exists by using os.Getwd()
// render index.html with go html/template file

func WriteDoc(Documents []interface{}, ApiDocGoVersion string) {
	tpl, err := template.New("index.gohtml").Funcs(template.FuncMap{
		"Time": func() string {
			return time.Now().Format("2006-1-2 15:4:5")
		},
		"Version": func() string {
			return ApiDocGoVersion
		},
		"Config": func() map[string]string {
			return DefaultConfig
		},
	}).ParseFiles("./template/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join("./assets/", "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = tpl.Execute(f, Documents)
	if err != nil {
		log.Fatal(err)
	}
}
