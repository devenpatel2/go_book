// demo for int datatype

package main

import "fmt"

/*
 int/unint are provided in 4 different sizes
  int8, int16, int32, int64. // uint8 uint16 uint32, uint64
  also there are the default int and uint which are either 32 or 64, 
  depending on the compiler's choice for a given hardwarde

  rune - synonym for int32 -  indicates that the value is Unicode
  byte - synonym for uint8 -  indiciates that the value is a peice of raw data
  uninptr - widht is not specified, but can hold all the bits of a pointer value

*/
 //  in case of overflow during an operation, higher bits are discareded
func main(){
    var u uint8 = 255
    fmt.Println(u, u+1, u*u) // "255 0 1"

    var i int8 = 127
    fmt.Println(i, i+1, i*i) // "127 -128 1"

    //bitwise operations
    // shift left with OR operation
    var x uint8 = 1<<1 | 1<<5
    var y uint8 = 1<<1 | 1<<2
    fmt.Printf("%08b\n", x)     // "00100010", the set {1, 5}
    fmt.Printf("%08b\n", y)     // "00000110", the set {1, 2}

    // AND operation
    fmt.Printf("%08b\n", x&y)    // "00000010", the intersection {1}
    //OR operation
    fmt.Printf("%08b\n", x|y)    // "00100110", the union {1, 2, 5}
    // XOR operation
    fmt.Printf("%08b\n", x^y)    // "00100100", the symmetric difference {2, 5}
    // AND NOT - bit clear operation
    fmt.Printf("%08b\n", x&^y)  // "00100000", the difference {5} . Clears bits of y, based on 1's in x

    for i := uint(0); i < 8; i++ {
        if x&(1<<i) !=0 { //membership test
            fmt.Println(i) // "1", "5"
        }
    }

    fmt.Printf("%08b\n", x<<1)   // "01000100", the set {2, 6}
    fmt.Printf("%08b\n", x>>1)   // "01000100", the set {0, 4}

    // binary operators must have operands of same type
    var apples int32 = 1
    var oranges int16 = 2
    // var compite int = apples + oranges // compiler error - mismatched types int32 and int16
    var compote = int(apples) + int(oranges)
    fmt.Printf("compote: %v\n", compote)

}
