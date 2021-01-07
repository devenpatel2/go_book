// uses recursion to print html node tree
// to run code 
// go run outline2.go http://gopl.io

package main
import (
    "fmt"
    "os"
    "net/http"
    "golang.org/x/net/html"
)

func main(){
    url := os.Args[1:][0]

    resp, err := http.Get(url)
    if err != nil{
        resp.Body.Close()
        fmt.Fprintf(os.Stderr, "outline: %v", err)
        os.Exit(1)
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        resp.Body.Close()
        fmt.Fprintf(os.Stderr, "outline: %v\n", err)
        os.Exit(1)
    }

    forEachNode(doc, startElement, endElement)
    resp.Body.Close()

}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil{
        pre(n)
    }

    // recursively traverse tree
    for c := n.FirstChild; c != nil ; c = c.NextSibling{
        forEachNode(c, pre, post)
    }

    if post != nil{
        post(n)
    }
}

var depth int

func startElement(n *html.Node){
    if n.Type == html.ElementNode {
        fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
        depth++
    }
}

func endElement(n *html.Node){
    if n.Type == html.ElementNode {
        fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
        depth--
    }
}
