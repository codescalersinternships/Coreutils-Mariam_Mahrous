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
	for i := 0; scanner.Scan() && i < n; i++ {
		fmt.Println(scanner.Text())
	}
	file.Close()
}
