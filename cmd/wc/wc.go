package main

import (
	"flag"
	"fmt"
	"os"
	s "strings"
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
	dat, err := os.ReadFile(s.Join(flag.Args(), " "))
	check(err)
	var line, word, char int = len(s.Split(string(dat), "\n")), 0, len(string(dat))
	var lastchar = ' '
	runes := []rune(string(dat))
	for _, grapheme := range runes {
		if grapheme != ' ' && grapheme != '\n' && (lastchar == ' ' || lastchar == '\n') {
			word++
		}
		lastchar = grapheme
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
	fmt.Printf("%s \n", s.Join(flag.Args(), " "))
}
