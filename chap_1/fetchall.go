// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main(){
    start := time.Now()
    // channel is a communication mechanism for goroutines
    // make channel of strings
    ch := make(chan string)
    for _, url := range os.Args[1:]{
        // goroutine is a concurrent function call
        // call fetch asynchronously
        go fetch(url, ch) //start a go routine
    }
    for range os.Args[1:]{
        // recieve all values on the channel
        fmt.Println(<-ch) // recieve from channel ch
    }
    fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    // reads body of response and discards it
    // io.Copy returns the byte count along with err
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() //don't leak resources
    secs := time.Since(start).Seconds()
    // send a line summary on the 'channel'
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
