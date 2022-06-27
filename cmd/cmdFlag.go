package cmd

import (
	"ApiDocGo/parser"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func ParseFlags() {
	var init bool
	flag.BoolVar(&init, "init", true, "initialization config.yml")
	flag.Parse()
	values := flag.Args()

	if len(values) == 0 {
		fmt.Println("Usage: apidocgo -i routes/ -o public/apidoc")
		flag.PrintDefaults()
		os.Exit(0)
	}

	for _, word := range values {

		if init {
			// fmt.Println(strings.ToUpper(word))
			strPath, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			data, err2 := yaml.Marshal(&parser.DefaultConfig)
			if err2 != nil {
				log.Fatal(err2)
			}
			err3 := ioutil.WriteFile(path.Join(strPath, "apidocgo.yml"), data, 0666)
			if err3 != nil {
				log.Fatal(err3)
			}

		} else {
			fmt.Println(word)
		}
	}
}
