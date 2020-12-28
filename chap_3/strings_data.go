// demo of strings data type

package main

import (
    "fmt"
    "unicode/utf8"
)

func main(){

    s := "hello, world"
    // built-in len returns numer of bytes
    // s[i] return the ith byte of string
    // for runes which encode utf-8, encoding non-ASCII char needs more than two bytes
    fmt.Println(len(s))      // 12
    fmt.Println(s[0], s[7])  // 104 119 ('h' and 'w')

    // attempting to access elements outside len range results in panic
    // c := s[len(s)]          // panic :  index out of range

    // substring s[i:j] yields a new string from i to j-1
    fmt.Println(s[0:5])        // hello
    // if any of i or j is ommited , the default values are 0 and len(s)
    fmt.Println(s[:5])         // hello
    fmt.Println(s[7:])         // world
    fmt.Println(s[:])          // hello, world

    // '+' operator makes a new string by concatenation
    fmt.Println("goodbye" + s[5:])
    // strings are immutables
    p := "left foot"
    t := p
    // p is modified to hold a new values
    // t however, still holds the old value
    p += ", right foot"
    fmt.Println(p)      // left foot, right foot
    fmt.Println(t)      // left foot

    // rune is a unicode code point
    // rune type is int32. so, int32 for purpose of unicode is rune
    // i.e. a sequence of runes is a sequence of int32
    // UTF-8 is variable lenght encoding of Unicode code points (rune)
    // It uses between 1-4 gytes to represnt a rune
    // 1 -byte for ASCII characters, 2-3 for most common rune
    // higher order 0 indicates 1 byte code, with following 7-bits are ASCII
    // higher order 110 indicates rune takes 2 bytes. the second byte begins with 10
    // go source files are always encoded in UTF-8

    // the two forms are - \uhhhh (16-bit value) and \uhhhhhhhh (32-bit value)
    // where h is the hex digit
    // all the string literals represent the same 6-byte string
    s1 := "世界"
    s2 := "\xe4\xb8\x96\xe7\x95\x8c"
    s3 := "\u4e16\u754c"
    s4 := "\U00004e16\U0000754c"
    fmt.Printf("%s\n%s\n%s\n%s\n",s1, s2, s3, s4)
    // length of string is not the same as rune-count or charcter count
    fmt.Printf("len(%s): %d\n", s1, len(s1))
    fmt.Printf("rune length: %d\n", utf8.RuneCountInString(s1))

    // to process the characters, utf decoder is needed
    // for each non-ascii rune, more than 1 byte is needed
    s5 := "Hello, 世界"
    for i := 0; i < len(s5); {
        // returns the rune r and the size of the rune
        r, size := utf8.DecodeRuneInString(s5[i:])
        fmt.Printf("%d\t%c\n", i, r)
        // update loop index by size
        i += size
    }
    // when range is applied to string, ut performs UTF-8 decoding implicitly
    for i, r := range "Hello, 世界" {
        fmt.Printf("%d\t%q\t%d\n", i, r, r)
    }
}
