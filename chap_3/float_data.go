package main
import (
    "fmt"
    "math"
)

func main(){
    var f float32 = 16777216 // 1 << 24
    fmt.Println(f == f+1)    // "true"!
    //largest numbers
    fmt.Printf("largest float32: %g\n", math.MaxFloat32)
    fmt.Printf("largest float64: %g\n", math.MaxFloat64)
    // formatted printing of floats
    for x := 0; x < 8; x++ {
        fmt.Printf("x = %d e A = %8.3f\n", x, math.Exp(float64(x)))
    }

    //infs and NaNs
    var z float64
    fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

    // any comparison with NaN ALWAYS yields false
    nan := math.NaN()
    fmt.Println(nan == nan, nan < nan, nan > nan)  // "false false false"
}
