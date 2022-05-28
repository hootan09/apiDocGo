package parser

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

var DocumentsArray [][]string
var Result []interface{}

// Reading definded path directory for finding go files & Documents
func ReadDoc(codesPath string) []interface{} {
	err := filepath.Walk(codesPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		// skip file with suffix _test.go
		if strings.HasSuffix(strings.ToLower(path), "_test.go") || filepath.Ext(path) != ".go" || info.IsDir() {
			return nil
		}

		// get documents as array
		commentsArray := GetTrimedComments(path)
		DocumentsArray = append(DocumentsArray, commentsArray...)

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", codesPath, err)
		return nil
	}

	// parsing Document array by parser
	ParseDocArray(DocumentsArray)
	return Result
}

// using go token parser for finding comments
func GetTrimedComments(path string) [][]string {

	var trimmedComments [][]string

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, path, nil, parser.ParseComments)

	if err != nil {
		log.Fatal("Error in parsing Comments ", err)
	}

	for _, c := range f.Comments {

		commentBlock := c.Text()

		// remove "*" from start line of comments if exists
		if strings.HasPrefix(commentBlock, "*") {
			commentBlock = strings.Replace(commentBlock, "*\n", "", 1)
			commentBlock = strings.Replace(commentBlock, "*", "", -1)
		}

		// split each comments base on @ (change to array of commentBlock)
		commentTag := strings.Split(commentBlock, "@api")

		// trim comments
		var newComments []string
		for _, val := range commentTag {
			if len(strings.TrimSpace(val)) != 0 {
				newComments = append(newComments, val)
			}
		}

		trimmedComments = append(trimmedComments, newComments)

	}
	return trimmedComments
}

func ParseDocArray(docArray [][]string) {
	for _, eachDoc := range docArray {
		var ParsedDoc = make(map[string]interface{})
		for _, line := range eachDoc {
			splitEachCommentBySpace := strings.SplitN(line, " ", 2)
			switch strings.ToLower(splitEachCommentBySpace[0]) {
			case "":
				fmt.Println(splitEachCommentBySpace[1])
				result := Api(splitEachCommentBySpace[1])
				ParsedDoc["api"] = *result
			case "name":
				result := Api_name(splitEachCommentBySpace[1])
				ParsedDoc["apiname"] = *result
			case "group":
				result := Api_group(splitEachCommentBySpace[1])
				ParsedDoc["apigroup"] = *result
			case "description":
				result := Api_group(splitEachCommentBySpace[1])
				ParsedDoc["apidescription"] = *result
			case "samplerequest":
				result := Api_sampleRequest(splitEachCommentBySpace[1])
				ParsedDoc["apisamplerequest"] = result
			case "version":
				result := Api_version(splitEachCommentBySpace[1])
				ParsedDoc["apiversion"] = result
			case "define":
				result := Api_define(splitEachCommentBySpace[1])
				ParsedDoc["apidefine"] = *result
			case "permission":
				result := Api_permission(splitEachCommentBySpace[1])
				ParsedDoc["apipermission"] = result
			case "header":
				result := Api_header(splitEachCommentBySpace[1])
				ParsedDoc["apiheader"] = *result
			case "headerexample":
				result := Api_headerExample(splitEachCommentBySpace[1])
				ParsedDoc["apiheaderexample"] = *result
			case "param":
				result := Api_param(splitEachCommentBySpace[1])
				ParsedDoc["apiparam"] = *result
			case "example":
				result := Api_example(splitEachCommentBySpace[1])
				ParsedDoc["apiexample"] = *result
			case "success":
				result := Api_success(splitEachCommentBySpace[1])
				ParsedDoc["apisuccess"] = *result
			case "error":
				result := Api_error(splitEachCommentBySpace[1])
				ParsedDoc["apierror"] = *result
			case "errorexample":
				result := Api_errorExample(splitEachCommentBySpace[1])
				ParsedDoc["apierrorexample"] = *result
			case "body":
				result := Api_body(splitEachCommentBySpace[1])
				ParsedDoc["apibody"] = *result
			case "use":
				result := Api_use(splitEachCommentBySpace[1])
				ParsedDoc["apiuse"] = result
			case "successexample":
				result := Api_successExample(splitEachCommentBySpace[1])
				ParsedDoc["apisuccessexample"] = *result
			case "query":
				result := Api_query(splitEachCommentBySpace[1])
				ParsedDoc["apiquery"] = *result
			case "paramexample":
				result := Api_paramExample(splitEachCommentBySpace[1])
				ParsedDoc["apiparamexample"] = *result
			default:
				log.Printf("@%s 'No Case match to parse'\n", splitEachCommentBySpace[0])
			}
		}
		Result = append(Result, ParsedDoc)
	}
}
