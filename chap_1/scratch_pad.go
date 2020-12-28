package main

import "fmt"

func main(){
    dict := make(map[string]int)

    dict["foo"] = 0
    if val, ok := dict["foo"]; ok{
        fmt.Println(val, ok)
    }

    fmt.Println(dict)
}
