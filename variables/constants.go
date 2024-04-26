package variables

import (
	"fmt"
	"math"
)

/*
1. A constant is anything that doesn't change its value in layman terms.
2. In Go, constants are declared using the const keyword.
3. Constants can be of any type like int, float, string, etc.
4. Constants can be declared at the package level or at the function level.
5. An important point to note is that constants can only be assigned a value at the time of declaration.
6. Unlike variables, constants cannot be declared without a value.
7. A constant name can only start with a letter or an underscore. It can be followed by any number of letters, numbers or underscores after that.
*/

const x1 int = 10         // global to the package
const X2 string = "Hello" // global to the package and can be accessed from outside the package
// const x3 := 10 // This will throw an error as constants cannot be declared without a value and := is not available at the package level

func TestConstants() {
	fmt.Println()
	fmt.Println("Constants")
	constants()
	exampleOfWhyWeNeedUnnamedConstants()
	globalConstants()
}

func constants() {
	/* Go has a very strong type system that doesn’t allow
	implicit conversion between any of the types.
	Even with the same numeric types no operation
	is allowed without explicit conversion.
	For eg you cannot add a int32 and int64 value.
	To add those either int32 has to be explicitly converted
	to int64 or vice versa. However untyped constant have the
	flexibility of temporary escape from the GO’s type system and
	can be used in places where a type is expected. */
	// Thus, in Go constants are of two types:
	// 1. Untyped constants
	// 2. Typed constants

	// Typed constants
	// const x int = 10
	// const y string = "Hello"
	// const z float64 = 20.5

	// Untyped constants
	// const x1 = 10
	// const y1 = "Hello"
	// const z1 = 20.5

	// Default hidden types of constants:
	// Integers	       int
	// Floats          float64
	// Complex Numbers complex128
	// Strings         string
	// Booleans        bool
	// Characters      int32 or rune

	//Unanamed untyped constant
	fmt.Printf("Type: %T Value: %v\n", 123, 123)
	fmt.Printf("Type: %T Value: %v\n", "circle", "circle")
	fmt.Printf("Type: %T Value: %v\n", 5.6, 5.6)
	fmt.Printf("Type: %T Value: %v\n", true, true)
	fmt.Printf("Type: %T Value: %v\n", 'a', 'a') // rune
	fmt.Printf("Type: %T Value: %v\n", 3+5i, 3+5i)

	//Named untyped constant
	const a = 123      //Default hidden type is int
	const b = "circle" //Default hidden type is string
	const c = 5.6      //Default hidden type is float64
	const d = true     //Default hidden type is bool
	const e = 'a'      //Default hidden type is rune
	const f = 3 + 5i   //Default hidden type is complex128

	fmt.Println("")
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
	fmt.Printf("Type: %T Value: %v\n", e, e) // rune
	fmt.Printf("Type: %T Value: %v\n", f, f)
}

func exampleOfWhyWeNeedUnnamedConstants() {
	var f1 float32
	var f2 float64
	f1 = math.Pi
	f2 = math.Pi

	// Due to the untyped nature of math.Pi constant,
	// it can be assigned to a variable of type float32 as well as float64.
	// This is otherwise not possible in GO after type is fixed.
	// When we print the type of math.Pi , it prints the default type which is float64.
	// But when we assign it to a float32 variable, it gets converted to float32.
	// Depending upon use case an untyped constant can be assigned to a low precision type (float32) or a high precision type(float64).
	// This is the reason why untyped constants are useful.

	fmt.Printf("Type: %T Value: %v\n", math.Pi, math.Pi)
	fmt.Printf("Type: %T Value: %v\n", f1, f1)
	fmt.Printf("Type: %T Value: %v\n", f2, f2)
}

func globalConstants() {
	/* Like any other variable, a constant will be global
	within a package if it is declared at the top of a file
	outside the scope of any function.
	For example, in the below program name will be a global constant
	available throughout the main package in any function.
	Do note that the const name will not be available outside the main package.
	For it to be available outside the main package it has to
	start with a capital letter. */
	fmt.Println()
	fmt.Println("Global constant visible to package            :", x1)
	fmt.Println("Global constant visible to package and beyond :", X2)
}
