// print contents found on URL
// demonstration of fetch.go from chapter 1 using defer
package main
import(
    "fmt"
    "io"
    "net/http"
    "os"
    "path"
    "strings"
)


func fetch(url string)(filename string, n int64, err error){
    resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()

    local := path.Base(resp.Request.URL.Path)
    if local == "/" {
        local = "index.html"
    }
    f, err := os.Create(local)
    if err != nil {
        return "", 0, nil
    }
    n, err = io.Copy(f, resp.Body)
    // Close file, but prefer error from Copy, if any.
    // closing of file is not defered since errors on file copy 
    // are reported on closing. 
    if closeErr := f.Close(); err == nil {
        err = closeErr
    }

    return local, n, err
}


func main(){
    for _, url := range os.Args[1:]{
        if !strings.HasPrefix(url, "http"){
            url = "http://" + url
        }
        filename, _, err := fetch(url)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Println(filename)
    }
}
