// Issues prints a table of GitHub issues matching the search terms.
// to run 
// go run issues.go is:open json decoder
package main
import (
    "fmt"
    "log"
    "os"
    "github"
)

func main(){
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %.55s\n",
            item.Number, item.User.Login, item.Title)
	}   
	

/* 
Partial output

0 issues:
#3       havardl json.decoder.JSONDecodeError
#24231   mshanak Desk translation page bug ( json.decoder)
#71    mosugo414 json decoder issue
#995   elliottwi Fix race condition in json decoder
#76    tathastu8 Error while running due to json module
#2693  greenbitl Fields dinamic decoder
#83        pk044 Circe Encoder/Decoder derivation
#28    guestofho json decoder error
#2     brottobhm error json
*/