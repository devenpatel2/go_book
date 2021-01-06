// demo code for functions

package main
import (
    "fmt"
    "math"
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

}