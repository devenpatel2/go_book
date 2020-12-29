//Understanding arrays 
package main
import (
    "crypto/sha256"
    "fmt"
)

func main(){
    var a [3]int               // array of 3 integers
    fmt.Println(a[0])          // print first element of a
    fmt.Println(a[len(a)-1])   // print last element of a

    // Print indices and elements
    // by default all elements are initialized to the zero-type
    for i, v := range a {
        fmt.Printf("%d %d\n", i, v)
    }

    // Print only elements
    for _, v := range a {
        fmt.Printf("%d\n", v)
    }

    //initialise by array array literal
    var q [3]int = [3]int{1, 2, 3}
    var r [3]int = [3]int{1,2}
    fmt.Printf("array q - %v\nr[2]:%d\n", q, r[2])

    // if ellipsis "..." appears in place of length,
    // the array is determined by number of initializers
    q1 := [...]int{1, 2, 3, 4}
    // format convereter verb "T" prints type of the variable
    // Size of the array is part of the 'type'. Thus [3]int and [4]int are differnt types
    // Size should be able to be computed during program compilation
    fmt.Printf("%T\n", q1)

    // this syntax can be used also for key value pair

    type Currency int
    const (
        USD Currency =  iota // const generator
        EUR
        GBP
        RMB
    )
    // the ellipsis denote that array length would be decided by initializers
    // initializeas as a key value pair.
    // the array symbol is still a string array. The initialization 
    // here is done using a key value pair
    // syntax - var := [...]var-type{idx1:val, idx2: val2}
    symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
    fmt.Println(RMB, symbol[RMB]) // 3 ¥

    // Indices can appear in any order, and also can be ommited. If an index is ommitted, 
    // the value takes on the zero-type for the element type. 
    // the following defines an array r with 100 elements. all are zero, except for the last (99)
    // which as the value -1
    r1 := [...]int{99:-1}
    fmt.Println(r1)

    // Comparision
    // If an array's element type are comparable, then the arrays are comparable too
    // Recall that the size of the array is part of the array-type
    // hence [2]int and [3]int would not be comparable

    a1 := [2]int{1, 2}
    b1 := [...]int{1, 2}
    c1 := [2]int{1, 3}
    fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"
    // d := [3]int{1,2}
    // fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

    //pass by ref
    // create a 32 size byte array
    // byte array is uint8 array
    sha_x := sha256.Sum256([]byte("x"))
    fmt.Printf("%x\n", sha_x)  // print hex values
    // pass by ref
    zero2(&sha_x)
    fmt.Printf("%x\n", sha_x)
    // fmt.Println(sha_x)


}


// set elements of byte arr to zero
func zero1(ptr *[32]byte){
    for i := range ptr {
        ptr[i] = 0
    }
}

func zero2(ptr *[32]byte){
    *ptr = [32]byte{} // set all values to zero-type
}
