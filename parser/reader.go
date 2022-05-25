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

// Reading definded path directory for finding go file
func ReadPath(codesPath string) {
	err := filepath.Walk(codesPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		// skip file with suffix _test.go
		if strings.HasSuffix(strings.ToLower(path), "_test.go") || filepath.Ext(path) != ".go" || info.IsDir() {
			return nil
		}

		// fmt.Println(info.Name())
		commentsArray := trimComments(path)
		fmt.Println(commentsArray)

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", codesPath, err)
		return
	}
}

// using go token parser for finding comments
func trimComments(path string) [][]string {

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
