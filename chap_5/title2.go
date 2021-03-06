// this program fetches HTML document and prints it's title
// this code demonstrates 'defered' functions in go
// to run code
// go run title1.go http://gopl.io
// go run title1.go https://golang.org/doc/effective_go.html


package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

// THe resp.Body.Close() call has to be made multiple times.
// defer statement is used to with operations like open/close, connect/disconnect, lock/unlock operations
// The right place for a defer statement that releases a resource is
// immediately after the res ource has been successfully acquired.
// In this case, it would be after the resp object is created

func title(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }

    defer resp.Body.Close()
    // Check Content-type is HTML (eg., "text/html; charset=utf-8")
    ct := resp.Header.Get("Content-Type")
    if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
        // resp.Body.Close()    // No longer needed
        return fmt.Errorf("%s has type %s, not text/html", url, ct)
    }

    doc, err := html.Parse(resp.Body)
    // resp.Body.Close()   // No longer needed
    if err != nil {
        return fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    visitNode := func(n *html.Node){
        if n.Type == html.ElementNode && n.Data == "title" &&
            n.FirstChild != nil {
                fmt.Println(n.FirstChild.Data)
            }
    }

    forEachNode(doc, visitNode, nil)
    return nil
}


// function copied from outline2.go
// functions can be passed as params with the correct function-type (signature)
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


func main() {
    for _, arg := range os.Args[1:] {
        if err := title(arg); err != nil {
            fmt.Fprintf(os.Stderr, "title: %v\n", err)
        }
    }
}
