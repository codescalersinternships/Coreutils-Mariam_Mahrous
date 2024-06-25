package main

import (
	"fmt"
	"os"
)

// Not sure if this is correct tbh
func main() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
