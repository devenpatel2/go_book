// this code demonstrates 'variadic' function
// variadic fuctions are functions that can accept variable number of arguments

package main

import (
    "fmt"
    "os"
)

func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

func main(){

    fmt.Println(sum())          // 0
    fmt.Println(sum(3))         // 3
    fmt.Println(sum(1,2,3,4))   // 10
    // the caller allocats and array and passes a slice of entire array
    // the above call behaves the same way as below
    values := []int{1,2,3,4}
    fmt.Println(sum(values...)) //10

    // variadic functions are often useful for string formatting.
    linenum, name := 12, "count"
    errorf(linenum, "undefined: %s", name) // Line 12: undefined: count
}

func errorf(linenum int, format string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, "Line %d:", linenum)
    fmt.Fprintf(os.Stderr, format, args...)
    fmt.Fprintln(os.Stderr)
}
