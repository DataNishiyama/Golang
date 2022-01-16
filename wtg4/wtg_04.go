package main

import "fmt"

func main() {

	//Constants: Storing data values

	//Unchageable values: constant
	//(bool, number(int, float, complex) and String)

	// Convention: use caps for CONST
	// Specify type optional
	const PI = 3.14159

	const B = "Hello World"

	// Specify type included
	const C9 string = "Specify: Hello World"

	fmt.Println(PI)

	fmt.Println(B)

	fmt.Println(C9)

	// Typed and Untyped Constatnts
	var n int = 2
	var sum int = (n + 5) //untyped numeric constant 5 = typed

	fmt.Println(sum)

	//Const eval at compile time
	const C1 = 2 / 3

	//Underclared Error getNumber used as value
	//const C2 = getNumber()
	//const

	//Numeric Constants = No size or sign
	// Arbitrary precision, no overflow
	const neg = -1
	const pos = 3

	var result = neg + pos
	fmt.Println(result)
	//Type Cast Compile Error
	//fmt.Println(string(result))

	//Multiple Assignment (Untyped Const)
	const BEEF, TWO, C = "meat", 2, "veg"
	fmt.Println(BEEF, TWO, C)

	//Multiple Specify Typed Constants
	const MONDAY, TUESDAY int = 1, 2

	fmt.Println(MONDAY, TUESDAY)

	//Listing all elements of set = enumeration
	//const (
	//UNKNOWN = 0
	//FEMALE  = 1
	//MALE    = 2
	//)

	//iota first use gives 0

	type Gender int
	const (
		UNKNOWN = iota
		FEMALE  = iota
		MALE    = iota
	)

	fmt.Println(MALE)

}
