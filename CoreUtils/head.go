package main 

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e!=nil{
		panic(e)
	}
}

func main() {
	arg := os.Args
	n:=10
	dat,err :=os.ReadFile(arg[1])
	check(err)
	if len(arg)==4 && arg[2]=="-n"{
		n, err =strconv.Atoi(arg[3])
		check(err)
	}
	var i,spaces int = 0,0
	for {
		fmt.Print(string(dat[i]))
		if dat[i]==10 {
			spaces++
		}
		i++
		if spaces==n || i==len(string(dat)) {
			break
		}
	}
}