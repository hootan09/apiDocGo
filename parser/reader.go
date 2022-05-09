package parser

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

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
		fmt.Println(info.Name())
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", codesPath, err)
		return
	}
}

// using go token parser for finding comments
