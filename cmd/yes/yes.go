package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args
	req := "y"
	if len(arg) != 1 {
		req = strings.Join(os.Args[1:], " ")
	}
	for {
		fmt.Println(req)
	}

}
