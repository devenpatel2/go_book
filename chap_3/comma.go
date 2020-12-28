package main
import (
    "fmt"
    "os"
)

func main(){
    num_s := os.Args[1:][0]
    fmt.Println(comma(num_s))
}


// comma inserts commas in a non-negative decimal integer string
func comma(s string) string{
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}
