// complex numbers demo
package main
import "fmt"

// go provides complex numbers - complex64 and complex128
// whose com[ponents are float32 and float 64 respectively
// built-in real and imag functions extracts the 
// real and imaginary parts

func main(){
    // var x complex128 = complex(1, 2) // 1+2i
    // var y complex128 = complex(3, 4) // 3+4i
    // or 
    x := 1 + 2i
    y := 3 + 4i
    fmt.Println(x * y)              // "(-5+10i), 1*3 + 1*4i + 2i*3 + 2i*4i
    fmt.Println(real(x*y))          // "-5"
    fmt.Println(imag(x*y))         // "10"

    // if a floating-point literal or decimal integer literal is immediately followed by i
    // e.g 3.1415i or 2i, it becomes 'imaginary literal'
    fmt.Println(1i * 1i)        //"(-1 + 0i), i^2 = -1
}
