// Package links provide a link-extraction function

package links

import (
    "fmt"
    "net/http"

    "golang.org/x/net/html"
)

// Extract HTTP GET request to the specified URL, parses
// the response as HTML, and retursn the links in the HTML document.

func Extract(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }

    var links []string
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            // iterate element to get all links
            for _, a := range n.Attr {
                if a.Key != "href"{
                    continue
                }
                link, err := resp.Request.URL.Parse(a.Val)
                if err != nil {
                    continue // ignore bad URLs
                }
                links = append(links, link.String())
            }
        }
    }

    // recursion using anonymous functions
    // function from outline2.go
    // Since only pre-function is needed for extraction, nil is passed as post-function
    forEachNode(doc, visitNode, nil)
    return links, nil
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
