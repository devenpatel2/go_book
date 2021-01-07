// demo code for functions

package main
import (
    "fmt"
    "math"
    "strings"
)

// function syntax
// func name(paramter-list) (return-list){
//    body
// }

func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}

// different ways to declare function
func add(x int, y int) int      {return x + y}
func sub(x, y int) (z int)      {z = x -y; return} // in this case z is initialized to zero
func first(x int, _ int) int    {return x}
func zero(int, int) int         {return 0}
func square(n int) int          {return n * n}
func negative(n int) int        {return -n}
func product(m, n int) int          {return m * n}
func add1(r rune) rune { return r + 1 }

func main(){
    fmt.Println(hypot(3,4))         // 5
    fmt.Printf("%T\n", add)         // func(int, int) int
    fmt.Printf("%T\n", sub)         // func(int, int) int
    fmt.Printf("%T\n", first)       // func(int, int) int
    fmt.Printf("%T\n", zero)        // func(int, int) int

    // the type of a function can also be called its 'signature'
    // two funcs have same parameter types and same return type are said to have the same signature
    // GO doesnt have default params, or pass by argument names :(
    // arguments are passed by value i.e func recieves a copy of each argument
    // when pointers or some reference is passed, variables are said to be 'indirectly' referred

    // functions are first-class values
    // they have type and they maybe pe assigned to variables
    // or passed to or returned from functions
    // funtion type variable can be declared
    // zero value of function type is nil
    // function type variables can only be compared to nil
    var f func(int) int
    fmt.Printf("f == nil: %v\n", f==nil)        // True

    f = square
    fmt.Println(f(3))       // 9

    f = negative
    fmt.Println(f(3))           // -3
    fmt.Printf("%T\n", f)   // funct(int) int

    // f = product          // compile error: can't assign f(int, int) int to f(int) int

    // functions can be passed as parameters to other functions
    // string.Map applies a function to each character of a string
    // the add1 function is applied to all characters of the string "HAL-9000"
    fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"

}
