package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Can't Read file %s\n", e)
		os.Exit(1)
	}
}

func numberCheck(n int) {
	if n < 0 {
		n = 0
		log.Fatal(e)
		fmt.Println("Try again using a valid postive number")
		os.Exit(1)
	}
}

func main() {
	arg := os.Args
	var n int
	flag.IntVar(&n, "n", 10, "No of lines displayed")
	flag.Parse()
	numberCheck(n)
	dat, err := os.ReadFile(arg[1])
	if len(arg) == 4 && n != 10 {
		dat, err = os.ReadFile(arg[3])
	}
	check(err)
	spaces := 0
	var s = ""
	for j := len(string(dat)) - 1; j >= 0; j-- {
		if dat[j] == 10 {
			spaces++
		}
		if spaces == n {
			break
		}
		q := string(dat[j])
		s = q + s
	}
	fmt.Print(s)
}
