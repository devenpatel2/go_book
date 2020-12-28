// Echo4 prints it's command-line args
// to run
// go run echo4.go -n -s @ 1 2 3

package main
// package flag uses programs cmd-line args to set value of variable in the program
import(
    "flag"
    "fmt"
    "strings"
)

// create a new flag variable of type bool
// args - name, default value, help string
var n = flag.Bool("n", false, "omit trailing newline")

// create a new flag variable of type string 
// args - name, default value, help string
// both n and s are pointers to flag variables
var sep = flag.String("s", " ", "seperator")

func main(){
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}
