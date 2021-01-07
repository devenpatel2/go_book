// a demonstration of error handling in go
// to run
// go run wait.go http://www.some-url.com
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)



// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.

func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    // Usual practice of go functions is that errors are dealt before success
    // and success is not indented within an else-block , but follows the outer-level
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil  //success
        }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries))  // exponential back-off
    }

    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main(){
    // add prefix to log messages
    log.SetPrefix("wait: ")
    url := os.Args[1:][0]
    if err := WaitForServer(url); err != nil {
        // fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        log.Fatalf("Site is down: %v\n", err)
        os.Exit(1)
    }
}
