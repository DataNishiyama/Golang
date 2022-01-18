package main

import (
	"fmt"
	"math"
)

type sideA = float64
type sideB = float64
type sideC = float64

// pythaGorean theorem
func pythaGo(a sideA, b sideB) sideC {

	var length sideC
	length = sideC(math.Sqrt((a * a) + (b * b)))
	return length
}

func main() {

	var lengthA sideA = 3
	var lengthB sideB = 6

	// Function Call
	hypoT := pythaGo(lengthA, lengthB)

	// Solution Output
	fmt.Printf("Given sides %.2f and %.2f, the hypotenuse is: %.2f", lengthA, lengthB, hypoT)

}
