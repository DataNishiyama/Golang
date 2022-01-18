package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Elementary Types")

	//Boolean, Numeric, Character

	//var b bool = true //True or False
	//fmt.Printf("Value of b: %t\n", b)

	//int and float point:
	//int8, int16, int32, int64

	//unsigned int
	//uint8, uint16, unint32, uint64

	//Float
	//float32, float64

	//int fastest processing = math = float64 expected

	var a int
	var b int32

	a = 15
	//b = a+a //compile error
	b = b + 5

	fmt.Printf("Value of b: %d\n", b)
	fmt.Printf("Value of a: %d\n", a)

	var n int16 = 3
	var m int32 = 4
	fmt.Printf("Types: %d, %d", m, n)
	//m + n mismatch types

	//m = n compile error

	//Basic Math
	var p int16 = 1
	var q int16 = 2

	var sum int16 = p + q

	fmt.Printf("\n\nSum = %d\n", sum)

	//Complex Numbers:
	//re + im complex64, complex128

	// c = complex(re, im)

	//random int
	x := rand.Int()

	//random in range
	y := rand.Intn(8)

	fmt.Printf("a is: %d\n", x)
	fmt.Printf("y is: %d\n", y)

	//Char Type

	var ch byte = 'A'
	fmt.Println(ch)

	//Unicode Package

	//test letter, test digit, test whiteSpace
	//unicode.IsDigit() // return bool value

	//end Program
}
