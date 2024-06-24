package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strconv"
	s "strings"
)

var a int = math.MaxInt

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if len(s.Split(path, "/")) <= a && path != "./" {
		fmt.Println(" ", path)
	}
	return nil

}

func main() {
	arg := os.Args
	if len(arg) == 3 && arg[1] == "-l" {
		var err error
		a, err = strconv.Atoi(arg[2])
		check(err)
	}
	if arg[2] < "0" {
		a = 0
		fmt.Println("Try again using a valid postive number")
		os.Exit(0)
	}
	err := filepath.WalkDir("./", visit)
	check(err)
}
