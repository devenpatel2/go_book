// Print text of each line that appears more than
// once in stdin, preceded by  its count

// to run 
// echo "hello\ncat\nhello" |./dup1

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){

    // make creates an empty map
    // map holds a key/value pair with const-time retrieve
    // key can be any type which can be compared using "=="
    // the int values are defaulted to 0
    counts := make(map[string]int)

    // bufio help i/o effecient and convenient
    // Scanner reads input and breaks it into lines or words
    input := bufio.NewScanner(os.Stdin)

    // for can have optional else
    // this is executed when the for condition is false

    // Each call to input.Scan reads the next line and remove \n
    for input.Scan(){
        // this line is equivalent to 
        // line = input.Text()
        // counts[line] = counts[line] + 1
        counts[input.Text()]++
    }

    // print results
    // for loop on map runs over keys. 
    // when used with range, it returns key and value of map
    for key, value :=  range counts{
        // print only dups
        if value > 1{
            // Printf produces formatted output
            fmt.Printf("%d\t%s\n", value, key)
        }
    }

}
