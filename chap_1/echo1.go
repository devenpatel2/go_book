// Print command line args
package main
import (
    "fmt"
    "os"
)

func main(){
    // declare  VARiables s, sep of the type STRING
    var s, sep string
    // for is the only loop statement
    // no paranthesis in the for statement
    // := --> short variable declaration
    // for condition{
    //  ..
    // } 
    // behaves like a while loop
    for i := 0; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println("Number of args ", len(os.Args))
    fmt.Println(s)
}
