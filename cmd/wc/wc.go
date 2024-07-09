package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	dat, err := os.ReadFile(flag.Args()[0])
	check(err)
	var lines, word, char int = len(strings.Split(string(dat), "\n")), 0, len(string(dat))
	for _, line := range strings.Split(string(dat), "\n") {
		word += len(strings.Split(string(line), " "))
	}
	if showLines {
		fmt.Printf("%d ", lines)
	}
	if showWords {
		fmt.Printf("%d ", word)
	}
	if showChars {
		fmt.Printf("%d ", char)
	}
	if len(arg) == 2 {
		fmt.Printf("%d %d %d ", lines, word, char)
	}
	fmt.Printf("%s \n", strings.Join(flag.Args(), " "))
}
