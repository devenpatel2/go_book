// detect duplicate strings and not print subsequent occurance
// this program demostrates use of map as a set

// to run 
// go run dedup.go
// -- enter words on prompt
// if the word has not been entered before, it will be 'echoed'

package main
import (
    "fmt"
    "bufio"
    "os"
)

func main(){
    seen := make(map[string]bool)   // a set of strings
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        // if the input is not present in map
        // add input to map and print
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)  
    }
}
