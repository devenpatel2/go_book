// test pacakage tempconv
package main

import (
    "fmt"
    "tempconv"
)

func main() {
    fmt.Printf("Room Temp is %v\n", tempconv.RoomTempC)
    fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
}
