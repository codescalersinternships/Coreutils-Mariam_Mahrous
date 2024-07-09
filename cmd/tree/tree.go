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
var Startingpath = "."

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
	if len(strings.Split(path, "/")) <= a+len(strings.Split(Startingpath, "/")) && path != Startingpath {
		pathSplited := strings.Split(path, "/")
		printLength(len(pathSplited))
		fmt.Println(" ", pathSplited[len(pathSplited)-1])
	}
	return nil

}

func printLength(depth int) {
	for range depth - 1 {
		fmt.Print("  ")
	}
	fmt.Print("|__")

}

func main() {
	flag.IntVar(&a, "l", math.MaxInt, "Depth of Directories")
	flag.Parse()
	if len(flag.Args()) > 0 {
		Startingpath = flag.Args()[0]
	}
	err := filepath.WalkDir(Startingpath, visit)
	check(err)
}
