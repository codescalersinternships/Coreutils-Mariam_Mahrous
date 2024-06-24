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
	var i, spaces int = 0, 0
	for {
		fmt.Print(string(dat[i]))
		if dat[i] == 10 {
			spaces++
		}
		i++
		if spaces == n || i == len(string(dat)) {
			break
		}
	}
}
