package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	s "strings"
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
	dat, err := os.ReadFile(s.Join(flag.Args(), " "))
	check(err)
	lines := s.Split(string(dat), "\n")
	for i := 1; i<=len(lines); i++ {
		if showLines || i == 1 {
			fmt.Printf("%d ", i)
		}
		fmt.Println(lines[i-1])
	}
}
