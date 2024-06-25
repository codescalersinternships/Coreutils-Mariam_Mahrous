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
	var n int
	flag.IntVar(&n, "n", 10, "No of lines displayed")
	flag.Parse()
	dat, err := os.ReadFile(s.Join(flag.Args(), " "))
	check(err)
	lines := s.Split(string(dat), "\n")
	for i, line := range lines {
		if i >= n {
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
