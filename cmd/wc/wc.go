package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Can't Read file %s\n", e)
		os.Exit(1)
	}
}

func main() {
	arg := os.Args
	var showLines, showWords, showChars bool
	flag.BoolVar(&showLines, "l", false, "Show Lines")
	flag.BoolVar(&showWords, "w", false, "Show Words")
	flag.BoolVar(&showChars, "c", false, "Show Chars")
	flag.Parse()
	var req = arg[1]
	if len(arg) > 1 {
		req = arg[len(arg)-1]
	}
	dat, err := os.ReadFile(req)
	check(err)
	var line, word, char int = 0, 0, len(string(dat))
	var lastchar = ' '
	runes := []rune(string(dat))
	for i := 0; i < len(string(runes)); i++ {
		if runes[i] == '\n' {
			line++
		} else if runes[i] != ' ' && runes[i] != '\n' && (lastchar == ' ' || lastchar == '\n') {
			word++
		}
		lastchar = runes[i]
	}
	if showLines {
		fmt.Printf("%d ", line)
	}
	if showWords {
		fmt.Printf("%d ", word)
	}
	if showChars {
		fmt.Printf("%d ", char)
	}
	if len(arg) == 2 {
		fmt.Printf("%d %d %d ", line, word, char)
	}
	fmt.Printf("%s \n", req)
}
