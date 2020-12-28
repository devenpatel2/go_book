// print command line args
package main
import (
    "fmt"
    "os"
    "strings"
)

func main() {

    // Types of declaration
    // s := "" - short variable declaration, used only within functions
    //           and not package-level variables
    //var s string - relies on default intialization
    //var s = "" useful for declaring multiple variables
    //var s string = "" explcitly mention variable type, which is redundant when default values
    //                  is same as assigned value
    s, sep := "", ""

    // range produces pair of values
    // index and value of the element at that index
    // go does not allow unused local varaiables
    // and hence a blank indentifier is used
    for _, arg := range os.Args[1:]{
        // the following operation is costly as 
        // for every iteration a new string is assigned to s
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
    // join strings without creating new strings for loop
    fmt.Println(strings.Join(os.Args[1:], " "))
    // or simple print the args wihtout worrying about formatting
    fmt.Println(os.Args[1:])
}
