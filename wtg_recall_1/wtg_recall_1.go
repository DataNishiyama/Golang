//Reserved Keywords

//Identifiers

//Case Sensitivity, no start w/ digit, no operators

//Pre-Dec Identifiers

//Ex: IOTA

//Blank Identifier

//Anonymous Identifier

//Program Basics

//Keywords, constants, variables, operators, types, functions]

//Delimeters and Punctuation

//Parenth(), Braces{}, Brackets[]

//Statements, no semi colon necessary

//Multiple statements can use semi-c, but not rec

//Library, module, or package. Go file = 1 package

//First line = file belongs to which package

//package name = lower case letters

//stand alone executable = main
//each app contains one main
package main

import "fmt"

//compile source file with package name other, main
//pack1 = object file is stored in pack1.a.

//each directory contains one package (Compile Order)

//link set of packages together with import

//fmt = needs functions from package fmt: formatted IO
//package names enclosed within double quotes

//Import loads declarations, not insert source code

//import two line, single line with semi-c
//import "fmt"; import "os"

//factor keyword
//import (

//"fmt"
//"os"
//)

//Factoring= calling keyword on multiple instances
//import called twice

//Visibility: Upper case identifier "Group1"
//Object is visible in code outside the package

//Lowercase identifiers not visible outside the package
//Private identifiers/variables

//Call with dot notation pack1.Object (Object exported)

//pack1.Object pack2.Object

//Alias

//import fm "fmt" fm = Alias

//No unnecessary code = build error

//Functions

//func functionName()

//main function = no parameters and no return

//If func return object
//fucn func_Name(param type1, param type2) ret type1{}

//Return Multi
//func func_Name (param1,...) (ret1 type1, ret2 type2){}

//Small Functions
func main() {

	fmt.Println("This is a small IO function")

}

//useMixedCaps to write multiWordNames
