package main 
	
import (
    "fmt"
    "os"
    "strings"
)

//Not sure if this is correct tbh 
func main()  {
	for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}
