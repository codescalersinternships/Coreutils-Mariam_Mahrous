package main

import (
	"bufio"
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
	file, err := os.Open(s.Join(flag.Args(), " "))
	check(err)
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if showLines || i == 1 {
			fmt.Printf("%d ", i)
		}
		fmt.Println(scanner.Text())
	}
}
