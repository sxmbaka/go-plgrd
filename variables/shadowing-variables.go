package variables

/*
A variable declared within an inner scope having the same name as
variable declared in the outer scope will shadow the variable in the outer scope.
This is called variable shadowing.
*/

import "fmt"

var x = 10

func ShadowingVariables() {
	fmt.Println()
	fmt.Println("Shadowing Variables")
	fmt.Println("package's i:", x)

	// package's i is being shadowed:
	x := 5
	// i above is a new variable.

	fmt.Println("main block:", x)

	// a new scope begins
	{
		// i is a new variable
		x := 6
		fmt.Println("inner block:", x)
	}
	// the scope ends

	// main's i
	fmt.Println("this function's i:", x)
}
