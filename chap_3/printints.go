// intsToString is like fmt.Sprintf(values) but adds commas.
package main
import (
    "bytes"
    "fmt"
)

func main(){
    fmt.Println(intsToString([]int{1,2,3})) // [1, 2, 3]
}

func intsToString(values []int) string {
    var buf bytes.Buffer
    // use WriteRune when appending UTF-8 encoding
    // for ASCII, WriteByte is fine
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}
