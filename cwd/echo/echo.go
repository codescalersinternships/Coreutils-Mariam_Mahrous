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
	if arg[1] == "-n" {
		fmt.Print(arg[2])
	} else {
		fmt.Println(arg[1])
	}

}
