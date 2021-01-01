// append function specialized for int slices

package main
import "fmt"

func main(){
    // declare slices x and y
    var x, y []int
    for i := 0; i < 10 ; i++ {
        // append i to slice x and assign it to y
        // if the cap(x) = len(x), the append with make a new slice
        // to accomodate the new element. 
        // since it's not known when this happens, it's a usual practice to
        // assing the same variable the result of an append i.e
        // x = append(x, i)
        // the resultant slice may or may not refer to the same underlying array
        y = appendInt(x, i)
        fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
        // update x with slice y
        x = y
    }
}

// appends and int to x and returns the appended slice
func appendInt(x []int, y int) []int{
    var z []int
    zlen := len(x) + 1
    // check if existing array has capacity to append an element 
    if zlen <= cap(x){
        // There is room to grow. Extend the slice
        // z and x share the same underlying array
        z = x[:zlen]
    } else {
        // There is insufficient space. Allocate new array.
        // Grow by doubling, for amortized linear complexity
        zcap := zlen
        if zcap < 2*len(x){
            zcap = 2*len(x)
        }
        // make new slice big enough to hold the result
        z = make([]int, zlen, zcap)
        // in this case, z and x no longer referes different underlying arrays
        // copy elements of once slice to another
        // copy returns the number of elements copied
        copy(z, x)
    }
    // copy the element y into the new space
    z[len(x)] = y
    // return slice
    return z
}
