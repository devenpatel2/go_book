// Package tempconv performs Celsius and Fahrenheit conversions.
// to initialise module
// go mod init tempconv
package tempconv

import "fmt"

// type declaration
// type name underlying-type
type Celsius float64
type Fahrenheit float64

const(
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

var RoomTempC Celsius
var RoomTempF Fahrenheit

// assiociate function 'String' to parameter c. 
// func (parameter param-type) func_name return_type {func body}
func (c Celsius) String() string { return fmt.Sprintf("%g C", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%g F", f)}
