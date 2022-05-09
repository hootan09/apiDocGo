package parser

// if JSON config file not exist must using this
var defaultConfig = map[string]string{
	"name":        "Acme project",
	"version":     "0.0.0",
	"description": "REST Api",
}

//TODO
//get apidocgo json config file if exists by using os.Getwd()
// render index.html with go html/template file
