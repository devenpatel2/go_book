//demo for const data type
package main
import (
    "fmt"
    "time"
)


func main(){
    // a const declaration may specify a  type. 
    // but in absence of type, it the type is inferred from the assigned value

    // time.Duration is named type whose underlying type is int64
    // time.Minute is a const of the type time.Duration
    const noDelay time.Duration = 0
    const timeout = 5 * time.Minute
    // %T print the data type of the varaible
    // [1] instructs to use the same variable for printing
    fmt.Printf("%T %[1]v\n", noDelay)       // time.Duration 0
    fmt.Printf("%T %[1]v\n", timeout)       // time.Duration 5m0s
    fmt.Printf("%T %[1]v\n", time.Minute)   // time.Duration 1m0s

    // when declaring group of const, the value of all but first can be omitted. 
    // in this case, the varaibles take the value and type of the previous variable
    // this feature is useful when used with iota
    const (
        a =1
        b
        c = 2
        d
    )
    fmt.Println(a, b, c, d) // 1 1 2 2

    // iota - const generator
    // the const generator is used to create sequence of related values
    // without the need of explicitly setting each value

    // set a named type for int
    type Weekday int
    const (
        Sunday Weekday = iota  // sets Sunday to zero
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )
    fmt.Printf("Sunday    %d\n" +
               "Monday    %d\n" +
               "Tuesday   %d\n" +
               "Wednesday %d\n" +
               "Thursday  %d\n" +
               "Friday    %d\n" +
               "Saturday  %d\n",
               Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
