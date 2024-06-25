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

func numberCheck(n int) {
	if n < 0 {
		log.Fatalf("Try again using a valid postive number \n")
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 10, "No of lines displayed")
	flag.Parse()
	numberCheck(n)
	dat, err := os.ReadFile(s.Join(flag.Args(), " "))
	check(err)
	lines:= s.Split(string(dat), "\n")
	j := len(lines) - n
	if len(lines)-n < 0 {
		j = 0
	}
	for i := j; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
