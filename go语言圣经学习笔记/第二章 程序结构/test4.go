package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c / 2)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius(f * 2)
}
func main() {
	var ccc Celsius = 1
	xx := CToF(ccc)
	fmt.Println(xx, ccc)
}
