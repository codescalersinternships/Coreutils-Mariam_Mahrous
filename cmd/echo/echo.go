package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var noNewLine bool
	flag.BoolVar(&noNewLine, "n", false, "Omit new line")
	flag.Parse()
	args := strings.Join(flag.Args(), " ")
	if noNewLine {
		fmt.Print(args)
	} else {
		fmt.Println(args)
	}

}
