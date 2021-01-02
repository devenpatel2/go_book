// demonstration of nested struct
// A struct type S cannot declare a field of the same type S. 
// i.e an aggregate value cannot contain itself
// However, S may declare a field of the pointer type *S
// i.e recursive nesting of structs is not allowed, 
// but one can declare a pointer to the same type

// binary tree for insertion sort
package main
import "fmt"

// binary tree structure
type tree struct{
    value       int
    left, right *tree
}

// sort values in places
func Sort(values []int) {
    var root *tree
    for _, v := range values {
        root = add(root, v)
    }
    appendValues(values[:0], root)
}


// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) [] int{

    if t != nil {
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
    }
    return values
}

func add(t *tree, value int) *tree {
    if t == nil{
        // Equivalent to return &tree{value: value}.
        t = new(tree)
        t.value = value
        return t
    }

    if value < t.value {
        t.left = add(t.left, value)
    } else {
        t.right = add(t.right, value)
    }
    return t
}

func main(){
    arr := []int{3,2,6,1,3,6,8,7}
    fmt.Println(arr)
    Sort(arr)
    fmt.Println(arr)
}
