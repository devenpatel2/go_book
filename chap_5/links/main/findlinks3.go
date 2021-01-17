
// BFS
// caller provides a worklist items to visit and a function vaalue f to call for each item
// to run
// go run findlinks3.go https://golang.org

// The pro cess end s when all reach able web pages have been craw led 
// or the memory of the computer is exhausted.


package main

import (
    "fmt"
    "log"
    "os"

    "../links"
)
// Breadth First Search (BFS) calls f for each item in the worklist.
// ANy items returned by f are added to the worklist.
// f is called at most once of each item.

func breadthFirst(f func(item string) []string, worklist [] string) {
    // dictionary to keep track of nodes that are already visited
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item]{
                seen[item] = true
                // the ... opeartor does unpacking of list returned by f
                // this causes all the items in the ;ist returned by f to
                // be appended to worklist
                worklist = append(worklist, f(item)...)
            }
        }
    }
}


// in the crawlers, items are URUs. this func is supplied to BFS
// this  will extract the urls, and return them, so that those URLs in turn are also visited

func crawl(url string) []string{
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil{
        log.Print(err)
    }
    return list

}

func main(){
    // Crawl the web breadth-first
    // starting from the command-line arguments
    breadthFirst(crawl, os.Args[1:])
}
