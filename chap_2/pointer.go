package main
import "fmt"

func main(){
    x := 1
    // assign address of x to a variable
    // p points to a variable of the type int
    // zeor-value of pointer is nil
    p1 := &x            // p, of type *int, points to x
    fmt.Println(*p1)    //"1"
    *p1 = 2             // equivalent to x = 2
    fmt.Println(x)

    var p2  =  f()
    // value of address is still accessible
    // even though the variable is out of scope
    fmt.Println(*p2)
    // each call to the function returns a different pointer
    fmt.Println(f() == f()) //false

    v := 20
    incr(&v)
    fmt.Printf("Value after incr call %d\n",v)
    fmt.Printf("Value after incr call %d\n", incr(&v))
}

// function f return an interger pointer
func f() *int {
    v := 10
    return &v
}

// increment function.
// takes integer pointer as argument and returns int
func incr(p *int) int{
    *p++
    // return statement for print/assignment operations
    return *p
}
