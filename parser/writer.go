package parser

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// if yml config file not exist must using this
var DefaultConfig = map[string]string{
	"name":        "Api Doc Go project",
	"version":     "0.0.1",
	"description": "REST Api Doc",
	"url":         "http://localhost:8080/apidocgo",
}

//TODO
//get apidocgo json config file if exists by using os.Getwd()
// render index.html with go html/template file

func WriteDoc(documents Documents, ApiDocGoVersion string) {

	//check if apidocgo.yml exist or not
	wdPath, _ := os.Getwd()
	ymlPath := path.Join(wdPath, "apidocgo.yml")
	if _, ymlConfErr := os.Stat(ymlPath); ymlConfErr == nil {
		ymlFile, _ := ioutil.ReadFile(ymlPath)
		yaml.Unmarshal(ymlFile, &DefaultConfig)
	}

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
		"CreateId": func(text string) string {

			cleaned := strings.Replace(text, ":", "", -1)
			return strings.Join(strings.Split(strings.Title(cleaned), "/"), "")
		},
		"GetUrlMethod": func() string {
			return strings.SplitN(DefaultConfig["url"], ":", 2)[0]
		},
		"GetStaticPath": func() string {
			return strings.SplitN(DefaultConfig["url"], ":", 2)[1]
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

	err = tpl.Execute(f, documents)
	if err != nil {
		log.Fatal(err)
	}
}
