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
	file, err := os.Open(s.Join(flag.Args(), " "))
	check(err)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	j := len(lines) - n
	if len(lines)-n < 0 {
		j = 0
	}
	for i := j; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
