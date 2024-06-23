package main 

import (
	"fmt"
	"os"
	s "strings"
)

func check(e error) {
	if e!=nil{
		panic(e)
	}
}

func main() {
	arg := os.Args
	var req=""
	if len(arg)==3 {
		req=arg[2]
	}
	dat,err :=os.ReadFile(arg[1])
	check(err)
	var line,word,char int = 0,0, len(string(dat))
	var lastchar byte = 32
	for i:=0 ; i<len(string(dat)) ; i++{
		if dat[i]==10 {
			line++
		} else if dat[i]!=32 && dat[i]!=10 && (lastchar==32 || lastchar==10 ) { 
			word++
		}
		lastchar=dat[i]
	}
	if s.Contains(req, "l") {
		fmt.Printf("%d ", line)
	}
	if s.Contains(req, "w") {
		fmt.Printf("%d ", word)
	}
	if s.Contains(req, "c") {
		fmt.Printf("%d ", char)
	}
	if len(arg)==2 {
		fmt.Printf("%d %d %d ", line , word , char)
	}
	fmt.Printf("%s \n", arg[1])
}

