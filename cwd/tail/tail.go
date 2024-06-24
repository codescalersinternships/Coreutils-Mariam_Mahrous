package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arg := os.Args
	n := 10
	dat, err := os.ReadFile(arg[1])
	if len(arg) == 4 && arg[2] == "-n" {
		n, err = strconv.Atoi(arg[3])
		check(err)
		if arg[3] < "0" {
			n = 0
			fmt.Println("Try again using a valid postive number")
			os.Exit(0)
		}
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
