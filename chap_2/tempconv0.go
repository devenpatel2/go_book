// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main
import "fmt"

// type declaration defines a new named-type that has the same
// "underlying type" as an existing type.
// type name underlying-type
type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -275.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func main(){
    fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
    boilingF := CToF(BoilingC)
    fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
    //fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch

    var c Celsius
    var f Fahrenheit
    fmt.Println(c == 0)     //"true"
    fmt.Println(f >= 0)     // "true"

    // fmt.Println( c == f) // compile error; type mismatch
    fmt.Println(c == Celsius(f)) //"true" , due to type conversion

    c = FToC(212.0)
    fmt.Println(c.String())  //"100 C"
    fmt.Printf("%v\n", c)    //"100 C"; no need to call String explicitly
    fmt.Printf("%s\n", c)    //"100 C"
    fmt.Println(c)           //"100 C
    fmt.Printf("%g\n", c)    //"100" ; does not call String func
    fmt.Println(float64(c))  //"100"; does not call String func
} 

// convert C to F and convert the data type
// conversion of one type to another is allowed if both have the same
// underlying data type OR
// bothe are unnamed pointers pointing to variables with same underlying type
func CToF(c Celsius) Fahrenheit {return Fahrenheit(c * 9/5 + 32)}

func FToC(f Fahrenheit) Celsius {return Celsius((f - 32) * 5 / 9)}


// func (parameter param_type) func_name(args) return_type {func body}
// the func 'String' is associated with the parameter 'c'
func (c Celsius) String() string {
    return fmt.Sprintf("%g C", c)
}
