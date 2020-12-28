// Server1 is a minimal "echo" server.
package main

import(
    "fmt"
    "log"
    "net/http"
)

func main(){
    // connect handler function to incoming URL
    http.HandleFunc("/", handler) // request call handler
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of request URL.
// r - struct type http.Request
func handler(w http.ResponseWriter, r *http.Request){
    // send back response using fmt.Fprintf
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
