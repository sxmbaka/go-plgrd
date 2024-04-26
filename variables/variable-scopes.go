package variables

import "fmt"

var gx1 int = 10         // gx1 is a global variable which cannot be accessed from outside the package
var Gx2 string = "Hello" // Gx2 is a global variable which can be accessed from outside the package

// try to access gx1 from outside the package and see the error
// e.g., go to main.go and try to access gx1 using
// fmt.Println(variables.gx1)
// this will throw an error as gx1 is not accessible outside the package

// try to access Gx2 from outside the package and see the result
// e.g., go to main.go and try to access Gx2 using
// fmt.Println(variables.Gx2)

/*
A variable declaration can be done at the package level or a function level or a block level.
The scope of a variable defines from where that variable can be accessed and also the lifetime of the variable.
Golang variables can be divided into two categories based on the scope:
1. Local variables
2. Global variables
*/

func TestAllScopes() {
	localScope()
	globalScope()
}

/*
1. Local variables are variables which are defined within a block or a function level
2. An example of a block is a for loop or a range loop etc.
3. These variables are only be accessed from within their block or function
4. These variables only live till the end of the block or a function in which they are declared. After that, they are Garbage Collected.
5. A local once declared cannot be redeclared within the same block or function.
*/
func localScope() {
	fmt.Println()
	fmt.Println("Local Scope")
	var x1 int = 10
	fmt.Println(x1)
	{
		var x2 int = 20 // x2 is a local variable to this block
		fmt.Println(x1) // x1 is accessible here
		fmt.Println(x2) // x2 is accessible here
	}
	// fmt.Println(x2) // This will throw an error as x2 is not accessible outside the block
}

/*
1. A variable will be global within a package if it is declared at the top of a file outside the scope of any function or block.
2. If this variable name starts with a lowercase letter then it can be accessed from within the package which contains this variable definition.
3. If the variable name starts with an uppercase letter then it can be accessed from an outside different package other than which it is declared.
4. Global variable are available throughout the lifetime of a program.
*/
func globalScope() {
	fmt.Println()
	fmt.Println("Global Scope")
	fmt.Println(gx1)
	fmt.Println(Gx2)
}
