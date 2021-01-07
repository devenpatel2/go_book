// demonstration of returning multiple values from a function
// to run the code 
// go run findlinks2.go https://golang.org 

package main
import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main(){
    var firstUrl string
    for _, url := range os.Args[1:]{
        firstUrl = url
        links, err := findLinks(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
            continue
        }
        for _, link := range links {
            //fmt.Println(link)
            // demonstration of calling a multi-value return function 
            // in the return statement of another function
            findLinksLog(link)
        }
    }

    fmt.Println(CountWordsAndImages(firstUrl))
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.

// The func returns two values, and thus each return statement
// returns a pair of values
func findLinks(url string)([]string, error) {
    resp, err := http.Get(url)
    if err != nil{
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        // network and OS resources should be released explicitly. 
        resp.Body.Close()
		// add context information with error
		// Errorf formats an error message using fmt.Sprintf and returns a new error value
		// it helps to build discriptive errors by prefixing additional context information
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        // add context information with error
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }

    return visit(nil, doc), nil
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

// the result of a multi-valued call iteself can be returned from 
// a multi value callking function

func findLinksLog(url string)([]string, error){
    fmt.Printf("findLinks: %s\n", url)
    return findLinks(url)
}


// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
// naming the  variables in return list for multiple return values 
// is particularly useful for documentation and readblity

// named variables in the return list are initialized to their zero value. 
// this allows us to do a ``bare return'' i.e a return statement without and parameters
func CountWordsAndImages(url string) (words, images int, err error){
    resp, err := http.Get(url)
    if err != nil {
        return  	// bare return
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil{
        err = fmt.Errorf("parsing  HTML: %s", err)
        return 		// bare return
    }

    words, images = countWordsAndImages(doc)
    return
}


// copied from https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-5-functions.html

func countWordsAndImages(n *html.Node) (words, images int) {
    // content in <style> or <script> are ignored
    if n.Type == html.ElementNode {
        if n.Data == "style" || n.Data == "script" {
            return
        } else if n.Data == "img" {
            images++
        }
    } else if n.Type == html.TextNode {
        text := strings.TrimSpace(n.Data)
        for _, line := range strings.Split(text, "\n") {
            if line != "" {
                words += len(strings.Split(line, " "))
                //fmt.Printf("%s %q %d\n", line, strings.Split(line, " "), len(strings.Split(line, " ")))
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i := countWordsAndImages(c)
        words += w
        images += i
    }
    return
}
