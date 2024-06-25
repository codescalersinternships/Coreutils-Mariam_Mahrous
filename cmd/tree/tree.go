package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var a int = math.MaxInt

func check(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if len(strings.Split(path, "/")) <= a && path != "./" {
		pathSplited := strings.Split(path, "/")
		printLength(len(pathSplited))
		fmt.Println(" ", pathSplited[len(pathSplited) -1])
	}
	return nil

}

func printLength(depth int) {
	for range depth - 1 {
		fmt.Print("  ")
	}
	fmt.Print("|__")
}

func numberCheck(n int) {
	if n < 0 {
		n = 0
		fmt.Println("Try again using a valid postive number")
		os.Exit(1)
	}
}

func main() {
	flag.IntVar(&a, "l", math.MaxInt, "Depth of Directories")
	flag.Parse()
	numberCheck(a)
	err := filepath.WalkDir("./", visit)
	check(err)
}
