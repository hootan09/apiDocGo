package cmd

import (
	"ApiDocGo/parser"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var (
	Init     bool
	Port     int
	ShowHelp bool
)

func ParseFlags() (port int) {
	port = 0
	ShowHelp := flag.Bool("h", false, "Print Usage Help")
	Init := flag.Bool("init", false, "initialization apidocgo.yml")
	Port := flag.Int("p", 0, "Port Addres for Serve Builded Documents")
	flag.Parse()

	// values := flag.Args()
	// if len(values) == 0 {
	// 	fmt.Println("Usage: apidocgo -i routes/ -o public/apidoc")
	// 	flag.PrintDefaults()
	// 	os.Exit(0)
	// }

	if *ShowHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//Create apidocgo.yml Config
	if *Init {
		strPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		data, err2 := yaml.Marshal(&parser.DefaultConfig)
		if err2 != nil {
			log.Fatal(err2)
		}
		err3 := ioutil.WriteFile(path.Join(strPath, parser.ConfigFileName), data, 0666)
		if err3 != nil {
			log.Fatal(err3)
		}
	}
	//serve Builded Document as web Server (note: the "Port" is public access from "cmd" package)
	if *Port > 0 {
		port = *Port
	}
	return port
}
