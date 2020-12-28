// Server1 is a minimal "echo" server.
package main

import(
    "fmt"
    "log"
    "net/http"
    "strings"
    "sync"
)
var mu sync.Mutex
var count int

func main(){
    // connect handler function to incoming URL
    http.HandleFunc("/", handler) // request call handler
    http.HandleFunc("/verbose", verboseHandler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of request URL.
// r - struct type http.Request
func handler(w http.ResponseWriter, r *http.Request){
    // send back response using fmt.Fprintf
    mu.Lock()
    count++
    mu.Unlock()
    if strings.HasPrefix(r.URL.Path, "/verbose"){
        verboseHandler(w, r)
    }

    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}


//counter echoes number of calls so far
func counter(w http.ResponseWriter, r *http.Request){
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}

//verbose echo
func verboseHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header{
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    // go allows simple statement to preceed an if condition
    if err := r.ParseForm(); err != nil{
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}
