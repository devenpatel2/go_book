// demonstration of Anonymous Functions
// When a function literal is used to denote a function using the 'func' keyword
// without the function name, it is called Anonymous Function
//
// e.g strings.Map( func (r rune) rune {return rune + 1}, "ABCD")
// 
// In the above example, the function strings.Map take two arguments - a function and a string
// the expression in first argument is an example for anonymous function

package main
import "fmt"

// the function squares takes no input and reutrs a 
// function with the return type int
func squares() func () int {
    var x int
    // this is an anonymous function
    // importantly, even after it is returned, it as acess
    // to the variables of the enclosing function
    // the lifetime of var x is not define by it's scope
    // x is hidden in the returning function
    return func() int {
        x++
        return x*x
    }
}

func main(){
    f := squares()
    fmt.Println(f())  // 1
    fmt.Println(f())  // 4
    fmt.Println(f())  // 9
    fmt.Println(f())  // 16
}
