package main

import "fmt"

//global scope (Declared outside function)
var number int = 5
var decision bool = false

func main() {

	fmt.Println("Scope of Variables")

	//Global and Local Scope

	//Scope of variables inside function = local scope / param and return variables

	fmt.Println("Original bool:", decision)

	var decision bool = true

	fmt.Println("Original value:", number)

	number = 10

	fmt.Println("New value:", number)

	fmt.Println("bool value:", decision)

	//Printf

	//func Printf (format string, list of var to print)
	//%d integral value, %s format string, %v default format

	fmt.Printf("Original Value of number: %d\n", number)

	fmt.Printf("New value of number: %d\n", number)

	fmt.Printf("Value of decision: %t\n", decision)

	//Value types
	fmt.Printf("Memory in a computer: words 32 or 64 / hex")

	//Values of primitive types : int, float, bool, string = value types: value in memory

	//References and Pointers = pointing to value or position in memory

	//end Program
}
