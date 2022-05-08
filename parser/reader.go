package parser

var defaultConfig = map[string]string{
	"name":        "Acme project",
	"version":     "0.0.0",
	"description": "REST Api",
}

//TODO
// seperate reader and write
// skip file with suffix _test.go
// needed to use filepath.walk
// using go token parser for finding comments
// render index.html with go html/template file
