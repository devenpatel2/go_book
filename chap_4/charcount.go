// Charcount computes counts of Unicode characters.
// to run 
// cat book.pdf > go run charcount.go

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main(){
	counts := make(map[rune]int)    // counts of Unicode characters
	// since encoding lenghts ranges only from 1 to utf8.UTFmax (4)
	// array is more compact then map
	var utflen [utf8.UTFMax + 1]int  // count of lenghts of UTF-8 encodings
	invalid := 0					// count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	// use of for loop as while loop

	for {
		// performs utf decoding and returns 
		// - decoded rune, length in bytes and an error value
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil{
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		// if input is not a legal utf-8 encoding, 
		// the returned value is unincode.Replacement and len is 1
		if r == unicode.ReplacementChar && n == 1{
			invalid++
			continue
		}

		// update rune count
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts{
		fmt.Printf("%c\t%d\n",c , n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)		
		}
	}

	if invalid > 0{
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)		
	}
}