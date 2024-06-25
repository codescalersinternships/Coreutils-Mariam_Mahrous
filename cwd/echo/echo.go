package main

import (
	"flag"
	"fmt"
	s "strings"
)

func main() {
	var showLines bool
	flag.BoolVar(&showLines, "n", false, "Omit new line")
	flag.Parse()
	args := s.Join(flag.Args(), " ")
	if showLines {
		fmt.Print(args)
	} else {
		fmt.Println(args)
	}

}
