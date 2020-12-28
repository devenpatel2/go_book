// Alternative approach for finding duplicate
// Instead of 'streaming' mode, read the file at once and then process it

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main(){
    counts := make(map[string]int)
    for _, filename := range os.Args[1:]{
        // Readfile returns byte slice that must be converted to string
        // which is then split into lines
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n"){
            counts[line]++
        }
    }

    for line, n := range counts{
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
