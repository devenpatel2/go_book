// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
// input param is a string slice - strings
// ouput is also a string slice
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        // if current element is empty, skip
        // else copy the current element to the 
        // last non-empty index
        if s != "" {
            // copy the current element to the previous empty element
            strings[i] = s
            i++
        }
    }
    // input and output share the same underlying array
    return strings[:i]
}


// implementation using append func
func nonempty2(strings []string) []string {
    out := strings[:0] // zero len slice of the orignal 
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}


func main(){
    data := []string{"one", "", "three", "" , "five", "six"}
    fmt.Printf("%q\n", nonempty2(data)) // ["one" "three". "five", "six"]
    // underlying array is partly over-written, thus 
    // data = nonempty(data)
    // should be preferred for in-place operations
    fmt.Printf("%q\n", data)           // ["one" "three" "five" "six" "five" "six"]
}

