package main

import (
	"fmt"
)

// alias type
type Celsius float32
type Fahrenheit float32

// conversion function
func toFahrenheit(t Celsius) Fahrenheit {

	var temp Fahrenheit
	temp = Fahrenheit((t * 9 / 5) + 32)
	return temp
}

// main function
func main() {

	var tempCelsius Celsius = 100
	// Function Call
	tempF := toFahrenheit(tempCelsius)
	fmt.Printf("%f Celsius is equal to %f F", tempCelsius, tempF)

	//end Program
}
