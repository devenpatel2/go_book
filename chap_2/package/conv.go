// conversion functions

package tempconv

// init is called on package initialization
// init functions can't be called or referenced. 
// they are automatically executed with the program starts
func init(){
    RoomTempC = Celsius(25)
    RoomTempF = Fahrenheit(77)
}

// CToF converts Celsius to Fahrenheit.
func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {return Celsius((f - 32) * 5 / 9)}
