package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	arg := os.Args
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	req := "y"
	if len(arg) == 2 {
		req = arg[1]
	}
	for {
		fmt.Println(req)
	}

}
