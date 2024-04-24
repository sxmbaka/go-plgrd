package variables

import (
	"fmt"
)

func TypeConversion() {
	/* Go is a statically typed language, which means
	that the data type of a variable is known at
	compile time. Once a variable is declared, it
	cannot be assigned a value of another type. */

	/* The following code will give an error
	because the variable a is declared as an int
	but we are trying to assign a float value to it */
	// var a int = 10.5

	/* Thus we need to convert the value to the
	correct data type before assigning it to a variable */

	/* The process of converting the value of
	one data type to another data type
	is called type conversion */

	/* In Go, type conversion is done by
	specifying the data type in which
	we want to convert the value in parentheses
	before the value */

	/* The syntax for type conversion is:
	<variable_name> = <data_type>(<value>) */

	/* The expression T(v) converts the value v to the type T. */

	fmt.Println()
	fmt.Println("Type Conversion")
	var a int = 65

	// Convert int to float32
	var c float32 = float32(a)
	fmt.Println("Converted a to float32:", c)

	// Convert float32 to int
	var d int = int(c)
	fmt.Println("Converted c to int:", d)

	/*Rune is an alias for int32
	and is used to represent Unicode code points
	the following code will print the character
	corresponding to the Unicode code point 65
	which is 'A'*/
	var e rune = rune(a)
	fmt.Println("Converted a to string:", e)

	// Add an example of a type conversion that is not possible in Go
	var s string = "123"
	// var i int = int(s) // This type conversion is not possible in Go
	// fmt.Println("Converted s to int:", i)
	// This will give an error because you cannot convert a string to an int directly
	// You will have to convert the string to a rune first and then to an int
	var i int = int(rune(s[0]))
	// This will convert the first character of the string to a rune and then to an int
	fmt.Println("Converted s to int:", i)

	// some type conversions may result in loss of data
	// for example, converting a float to an int
	// will result in the loss of the decimal part
	var f float32 = 10.5
	var g int = int(f)
	fmt.Println("Converted f to int:", g)

	// converting an int to a float will not result in loss of data
	// because the float data type can store integer values as well
	var h int = 10
	var i1 float32 = float32(h)
	fmt.Println("Converted h to float32:", i1)
}
