package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	var showLines bool
	flag.BoolVar(&showLines, "n", false, "Show Lines")
	flag.Parse()
	if showLines {
		fmt.Print(arg[2])
	} else {
		fmt.Println(arg[1])
	}

}
