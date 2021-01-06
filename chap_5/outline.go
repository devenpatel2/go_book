// uses recursion to print html node tree
// to run code 
// curl -s https://golang.org | go run outline.go

package main
import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main(){
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "outline: %v\n", err)
        os.Exit(1)
    }
    outline(nil, doc)
}

// print treee structure of HTML page using recursion
// as it traverses, each element is pushed to a stack and printed. 
func outline(stack []string, n *html.Node){
    if n.Type == html.ElementNode {
        // the stack slice is received through a recursive call
        // this statment will append to the stack and could even
        // modify the underlying array. however, since initial elements
        // are not modified, the elements of the stack are intact, 
        // even though there is no corresponding pop 
        stack = append(stack, n.Data)    // push tag
        fmt.Println(stack)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        outline(stack, c)
    }
}