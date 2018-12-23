package tempconv

import (
	"fmt"
)

//CtoF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
fmt.Printf("Brrrr!%v\n",tempconv.AbsoluteZeroC)  //"Berrrr! -273.15。C"
fmt.Println(tempconv.CToF(tempconv.BoilingC))   //"212。F