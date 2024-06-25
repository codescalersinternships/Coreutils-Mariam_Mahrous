package main

import (
	"flag"
	"fmt"
	s "strings"
)

func main() {
	var noNewLine bool
	flag.BoolVar(&noNewLine, "n", false, "Omit new line")
	flag.Parse()
	args := s.Join(flag.Args(), " ")
	if noNewLine {
		fmt.Print(args)
	} else {
		fmt.Println(args)
	}

}
