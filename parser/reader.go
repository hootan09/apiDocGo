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
		commentTag := strings.Split(commentBlock, "@")

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
			splitEachCommentBySpace := strings.Split(line, " ")

			switch strings.ToLower(splitEachCommentBySpace[0]) {
			case "api":
				result := Api(strings.Join(splitEachCommentBySpace[1:], " "))
				ParsedDoc["api"] = *result
			case "apiname":
				result := Api_name(strings.Join(splitEachCommentBySpace[1:], " "))
				ParsedDoc["apiname"] = *result
			case "apigroup":
				result := Api_group(strings.Join(splitEachCommentBySpace[1:], " "))
				ParsedDoc["apigroup"] = *result
			case "apidescription":
				result := Api_group(strings.Join(splitEachCommentBySpace[1:], " "))
				ParsedDoc["apidescription"] = *result
			case "apisamplerequest":
				result := Api_sampleRequest(strings.Join(splitEachCommentBySpace[1:], " "))
				ParsedDoc["apisamplerequest"] = result
			case "apiversion":
				result := Api_version(strings.Join(splitEachCommentBySpace[1:], " "))
				ParsedDoc["apiversion"] = result
			default:
				log.Printf("@%s 'No Case match to parse'\n", splitEachCommentBySpace[0])
			}
		}
		Result = append(Result, ParsedDoc)
	}
}
