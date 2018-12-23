package tempconv

import (
	"fmt"
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {

	fmt.Printf("brrrr! %v\n", tempconv.AbsoluteZeroC)

	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)

	fmt.Println(tempconv.CToF(tempconv.BoilingC))

}
