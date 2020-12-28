// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

// to run
// echo "hello\ncat\nhello" |./dup2
// ./dup2 f1.txt f2.txt

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0{
        countLines(os.Stdin, counts)
    } else{
        for _, arg := range files{
            // os.Open returns 2 values - an open file, built-in error type
            f, err := os.Open(arg)
            // check if file was opened successfully
            if err != nil {
                // %v - any value in a natural format
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }

    for line, n := range counts{
        if n > 1{
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

// functions may be declared in any order
// count line function
// when map is passed to function, the function recieves a copy to the reference
func countLines(f *os.File, counts map[string]int){
    input := bufio.NewScanner(f)
    for input.Scan(){
        counts[input.Text()]++
    }
}
