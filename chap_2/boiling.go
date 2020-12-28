package main
import "fmt"

// package level declaration
// visible throughout the file and also all 
// files of the package

const boilingF = 212.0

func main(){
    // local variables
    // variable declaration 
    // var name type = expression
    // type or "= expresion" might be omitted, but not both
    // if expression is omitted, var is initialized by the zero value of it's type
    // string - ""
    // numbers - 0
    // bool - false
    // interfaces, ref types - nil
    // for aggregator type like array, all elements are initialized to the type's zero value
    // := is a declaration "=" is assignment
    // short variable declaration (:=) are usually used to declare and initialize majority of local variable
    // 'var' declaration tends to be reserved for local variables that needs an explicit type
    var f = boilingF
    const freezingF = 32.0
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g F or %g C\n", f, c)
    fmt.Printf("%g F = %g C\n", freezingF, fToC(freezingF))
}

// func definition
// func name(variable_name var_type) return_type{
func fToC(f float64) float64{
    return (f - 32) * 5 / 9
}
