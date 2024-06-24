package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Can't Read file %s\n", e)
		log.Fatal(e)
	}
}

func printLineNumber(i int, showLines bool) {
	if showLines {
		fmt.Printf("%d ", i)
	}
}

func main() {
	arg := os.Args
	var showLines bool
	flag.BoolVar(&showLines, "n", false, "Show Lines")
	flag.Parse()
	dat, err := os.ReadFile(arg[1])
	if showLines {
		dat, err = os.ReadFile(arg[2])
	}
	check(err)
	line := 1
	printLineNumber(line, showLines)
	for i := 0; i < len(string(dat)); i++ {
		fmt.Print(string(dat[i]))
		if dat[i] == 10 {
			line++
			printLineNumber(line, showLines)
		}
	}
	fmt.Printf(" \n")
}
