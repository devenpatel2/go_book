// basename removes directory components and a .suffix.
package main
import (
    "fmt"
    "os"
)

func main(){
    path := os.Args[1:][0]
    fmt.Println(basename(path))
}

// Discard last '/' and everything before
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
    for i := len(s) -1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }
    // Preserver everything before the last '.'
    for i := len(s) -1; i >= 0; i--{
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    return s
}
