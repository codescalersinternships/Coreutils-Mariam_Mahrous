package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Can't Read file %s\n", e)
	}
}

func incrementLines(showLines bool) func() {
	i := 1
	return func() {
		if showLines {
			i++
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	args := os.Args
	var showLines bool
	flag.BoolVar(&showLines, "n", false, "Show Lines Numbers")
	flag.Parse()
	dat, err := os.ReadFile(args[1])
	if showLines {
		dat, err = os.ReadFile(args[2])
	}
	check(err)
	nextInt := incrementLines(showLines)
	nextInt()
	for i := 0; i < len(string(dat)); i++ {
		fmt.Print(string(dat[i]))
		if dat[i] == 10 {
			nextInt()
		}
	}
	fmt.Printf(" \n")
}
