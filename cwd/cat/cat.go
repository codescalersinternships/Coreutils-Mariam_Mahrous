package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arg := os.Args
	dat, err := os.ReadFile(arg[1])
	check(err)
	line := 1
	fmt.Printf("%d ", line)
	for i := 0; i < len(string(dat)); i++ {
		fmt.Print(string(dat[i]))
		if dat[i] == 10 {
			line++
			fmt.Printf("%d ", line)
		}
	}
	fmt.Printf(" \n")
}
