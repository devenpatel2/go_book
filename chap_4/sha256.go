// demonstrate array comparision
package main
import (
    "fmt"
    "crypto/sha256"
)

func main(){
    //compute the sha for 'x' and 'X'
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    // Even though the two inputs differ only by a bit, half the bits are completely different
    // %x - print all elements of arry or slice of bytes in hex
    // %t - show boolean
    // %T - show type of value
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    // Output:
    // 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
    // 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
    // false
    // [32]uint8
}
