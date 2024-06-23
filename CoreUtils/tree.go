package main 

import (
	"fmt"
	"io/fs"
    "path/filepath"
)

func check(e error) {
	if e!=nil{
		panic(e)
	}
}

func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path)
    return nil
}

func main() {
	err := filepath.WalkDir("./", visit)
	check(err)
}