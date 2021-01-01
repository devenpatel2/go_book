// demo code for slices
package main
import "fmt"
// slices are variable lenght sequences
// syntax []T , where T is the data-type.
// looks like an array, but wihtout the size specification

// Slice has 3 components - 
//    pointer  - points to first element of the array reachable through slice
//    length   - number of slice elements
//    capacity - number of elements between start of slice and the end of underlying array
  
//
func main(){
	// declare and initialise a string array
	// month[1] - January
	// month[12] - December
	// month[0] - ""  (zero value for string)
	months := [...]string{1: "January",
						  2: "February",
						  3: "March",
						  4: "April",
						  5: "May",
						  6: "June",
						  7: "July",
						  8: "August",
						  9: "September",
						  10: "October",
						  11: "November",
						  12: "December"}
	
	// built-ins len and cap return length and capacity of slices
	// s[i:j] creates a new slice that refers to elements of sequence s from i to j-1
	// if i is omitted, i = 0
	// if j is omitted, j = len(s)
	fmt.Println(months[1:])    // [Januaray ... December]
	fmt.Println(months[1:13])  // [Januaray ... December]

	// define two slices for quarter-2 and summer
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Printf("Capacity of Q2 %d\n", cap(Q2))           // 9 - capacity is from start of slice to end of array
	fmt.Printf("Capacity of summer %d\n", cap(summer))   // 7

	// iterate slices and check of common elements
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	// slicing beyond cap(s) causes panic
	// fmt.Println(summer[:20]) // panic: out of range
	// slicing beyond len(s) extends the slice

	// even though len(summer) is 3, this works.
	endlessSummer := summer[:5]		// extend a slice (within capacity)
	fmt.Println(endlessSummer) // "[June July August September October]"


	//unlike arrays, slices are not comparable using "=="	
	fmt.Println(equal(Q2, summer))   // false
	// however, this "deep equivalance" is not always practical 
	// for e.g a slice could contain itself
	// s := []interface{}{"one", nil}
	// s[1] = s
	// s2 := s[1].([]interface{})
	// s3 := s2[1].([]interface{})
	// No matter how deep we go, the 2nd element will always be the slice value pointing to the same array as s, wrapped in an interface{} value.

	// slice elements are indirect
	// The only legal slice comparison is against nil
	
	if summer == nil {
		fmt.Println("summer is nil")
	}

	// zero value of slice is nil
	//  - nil slice has no underlying array
	//  - slice of of lenght and capacity 0 are could still be non-nil slice
	//    []int{} or make([]int, 3)[3:] are no-nil slices with len and cap 0	

	// to test if slice is empty, use len and not nil to compare
	var s []int 
	fmt.Printf("len(s) - %d, s == nil - %t\n",len(s), s == nil)  // len(s) - s == nil - true
	s = nil 
	fmt.Printf("len(s) - %d, s == nil - %t\n",len(s), s == nil)  // len(s) - s == nil - true
	s = []int(nil)
	fmt.Printf("len(s) - %d, s == nil - %t\n",len(s), s == nil)  // len(s) - s == nil - true
	s = []int{}
	fmt.Printf("len(s) - %d, s == nil - %t\n",len(s), s == nil)  // len(s) - s == nil - false

	// built-in function make can be used to create slices of a specified type, len and cap
	// make([]T, len)
	// make([]T, len, cap)  //same as make([]T, cap)[:len]
	make_s := make([]int, 3)
	fmt.Printf("slice of type %T. len(make_s) - %d , cap(make_s) - %d\n", make_s, len(make_s), cap(make_s)) // slice of type []int. len(make_s) - 3 , cap(make_s) - 3
	// len as to be smaller than capacity
	make_s2 := make([]int, 3 , 6)
	fmt.Printf("slice of type %T. len(make_s2) - %d , cap(make_s2) - %d\n", make_s2, len(make_s2), cap(make_s2)) // slice of type []int. len(make_s2) - 3 , cap(make_s2) - 3

	
	// built-in append function appends items to slices
	// the below code is equivalent declaring and initializing a slice
	// var runes []rune("Hello, 世界") .
	// it demonstrates use of append to build a slice of nine runes

	// declare an empty rune
	var runes []rune
	for _, r := range "Hello, 世界" {
		//iteratively append elements to rune
		// since the resultant slice may or may not refer to the same underlying array
		// it is usual practice to assing the result of append to the same variable
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)   // ['H' 'e' 'l' 'l' 'o' ',' ' ' ' B ' ' F ']

	// updating the variable is required for any function that changes the length or capacity of the slice
	// since this may result in refering to a different underlying array

	var x []int
	x = append(x,1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	// varaible followed by ellipsis does slice unpacking
	// in-build append accepts any number of final arguments
	x = append(x, x...)    // append slice x
	fmt.Println(x)         // [1 2 3 4 5 6 1 2 3 4 5 6]


	// slices as as stack
	var stack []int

	for _, v := range x {
		stack = append(stack, v)        // push v
		fmt.Printf("Push %d, stack %d\n", v, stack)
	}
	

	for i := len(stack); i > 0 ; i-- {
		top := stack[len(stack) -1 ]      // top of stack
		stack = stack[:len(stack) - 1]   //shrkink stack
		fmt.Printf("Pop %d, stack %d\n", top, stack)
	}

	remove_test := []int{5,6,7,8,9}
	fmt.Println(remove(remove_test, 2))
}


// remove an element, preserving the order of elements
// 
func remove(slice []int, i int) []int {
	// copy(dst, src)
	copy(slice[i:], slice[i+1:])  // over writes the slice from i to end with slice from i+1 to end
	return slice[:len(slice)-1]
}


// one way to do compare equality, is to implement it ourselves
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}