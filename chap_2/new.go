//Use new function to create variables
package main
import "fmt"

func main(){
    // new(T) creates an unnamed variable of the type T
    p  := new(int)  // p, of the type *int, points to unnamed int variable
    fmt.Println(*p) // "0"
    *p = 2          // sets the unnamed int to 2
    fmt.Println(*p)
}
