// demostration of defer functionality

package main
import "fmt"

func double(x int) (result int) {
    // use anonymous function for defer
    // since the defer function is called 'after' the
    // returned values are updated, they can be captured by defered function
    defer func () { fmt.Printf("double (%d) = %d\n", x, result) } ()
    // the result variable is automaticall updated with the retured value
    return 2 * x
}


func triple(x int) (result int) {
    // defer function can also change the returned values !
    defer func () { result += x}()
    return double(x)
}

// silly example to demonstrate execution order of defered functions
// defer in a loop is generally not a good idea
// the function could be refactored so the the defered call is in a 
// sepearte function call
func defer_demo(){
    fmt.Println("Defer loop\n")
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d\n", i)
    }
}

func main(){
    _ = double(4)           // double(4) = 8
    fmt.Println(triple(4))  // 12
    defer_demo()
}
