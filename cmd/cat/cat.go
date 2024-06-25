package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Can't Read file %s\n", e)
	}
}

func main() {
	var showLines bool
	flag.BoolVar(&showLines, "n", false, "Show Lines Numbers")
	flag.Parse()
	dat, err := os.ReadFile(strings.Join(flag.Args(), " "))
	check(err)
	lines := strings.Split(string(dat), "\n")
	for i, line := range lines {
		if showLines || i == 0 {
			fmt.Printf("%d ", i+1)
		}
		fmt.Println(line)
	}
}
