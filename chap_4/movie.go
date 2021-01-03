// demonstration of JSON strings in Go
package main
import (
    "fmt"
    "encoding/json"
    "log"
)

// by default, the GO struct field names are chosen as the field names
// for JSON object. However using "reflection", this can be modified
type Movie struct {
    Title   string
    // optional struct field names for json
    // Year will be replaced by 'released' in json string
    // Color will be replacedd by "color" in json string
    // Also, if the value of Color is false, it will be skipped/omitted

    Year    int `json:"released"`    // field tag - string of meta-data

    // field tag is a string literal of key:value pair
    // the key 'json' controls behaviour of the encoding/json package
    // this is useful if JSON is using a differnt naming convention
    // i.e GO's PascalCase can be easily replaced to JSON prefrerred snake_case
    Color   bool `json:"color,omitempty"`

    Actors  []string
}

func main(){
    var movies = []Movie{
        {Title: "Casablanca", Year: 1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
        {Title: "Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Paul Newman"}},
        {Title: "Bullitt", Year: 1968, Color: true,
            Actors: []string{"Steve McQueen", "Jacqueline Bisset"}}, // comma is needed for the last element too
    }

    // to convert Go data structure to JSON is called `marshalling`
    // json.Marshal returs a byte slice and err status
    // Marshal however produces string that is difficult to read

    // data, err := json.Marshal(movies)
    // one can use MarhsalIndent to get a human readable format
    // additional args are prefix for each line and indentation level for each level
    data, err := json.MarshalIndent(movies, "", "  ")
    if err != nil {
        log.Fatalf("JSON marshalling failed %s\n", err)
    }
    fmt.Printf("%s\n", data)

    // decoding (or unmarshalling) of JSON string to GO structure can be acomplished by 
    // json.Unmarshal
    var movies_dec = []Movie{}
    if err := json.Unmarshal(data, &movies_dec); err != nil {
        log.Fatalf("JSON unmarshalling failed, %s\n", err)
    }
    fmt.Println(movies_dec)
    // [{Casablanca 1942 false [Humphrey Bogart Ingrid Bergman]} {Cool Hand Luke 1967 true [Paul Newman]} {Bullitt 1968 true [Steve McQueen Jacqueline Bisset]}]
    fmt.Println(movies_dec[0].Title)   // Casablanca

    // Unmarshalling also allows only a particular field to be extracted
    // declare an array of struct of type {Title string}
    var titles []struct{Title string}

    if err := json.Unmarshal(data, &titles); err !=  nil{
        log.Fatalf("JSON unmarshalling failed, %s\n", err)
    }
    fmt.Println(titles)     // [{Casablanca} {Cool Hand Luke} {Bullitt}]

}
