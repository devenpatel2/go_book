// Findlinks1 prints the links in an HTML document read from standard input.

// to run code 
// curl -s https://golang.org | go run findlinks1.go

package main
import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

// net/html is a non-standard html package that has useful functions for parsing html pages

// ref node structure for html package
/*
type NodeType int32

type Attribite struct {
    Key, Val String
}

type Node struct {
    Type                    NodeType
    Data                    string
    Attr                    []Attribute
    FirstChild, NextSibling *Node
}

func Parse(r io.Reader) (*Node, error)
*/

func main() {
    doc, err := html.Parse(os.Stdin)   // returns pointer to Node , error
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
        os.Exit(1)
    }

    for _, link := range visit(nil, doc){
        fmt.Println(link)
    }
}

// visit appends to links each link found in n and returns the result.
// it traverses the page recursilvely using the html attributes.

func visit( links []string, n *html.Node) []string {
    // look for anchor element <a href='...'>
    if n.Type == html.ElementNode && n.Data == "a"{

        for _, a := range n.Attr {
            // append anchor element if it is a hyper-link
            if a.Key == "href" {
                links = append(links, a.Val)
            }
        }
    }

    // traverse page recursively
    for c := n.FirstChild; c != nil; c=c.NextSibling{
        links = visit(links, c)
    }

    return links
}
