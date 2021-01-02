// demo for passing slice to functions

package main
import "fmt"

// since slice contains pointer to an element of array,
// passing slice allows to modify the underlying array elements

func main(){
    arr1 := [...]int{0,1,2,3,4,5}

    // the array is passed as whole. 
    // since the input of reverse is a slice, the array has to be passed as slice
    reverse(arr1[:])
    fmt.Println(arr1)   // [5 4 3 2 1 0]

    // arr2 is initialized as slice as opposed to arr1, which is initialized as array
    // This creates an array variable of right size and yields a slice
    arr2 := []int{0,1,2,3,4,5}
    // Rotate s left by two positions
    // reverse first to elements
    reverse(arr2[:2])
    // reverse last n-2 elements
    reverse(arr2[2:])
    // final reverse for rotation
    // since arr2 is a slice, it can be passed directly
    reverse(arr2)
    fmt.Println(arr2)    // [2 3 5 5 0 1]

}

// the input to this func is a slice
// notice the []int
// reverse reverses a slice of ints in place.
func reverse(s []int) {
    for i, j := 0, len(s) -1 ; i<j ; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
