package variables

import (
	"fmt"
)

func TestAllOperations() {
	integerOperations()
	floatingOperations()
	booleanOperations()
	complexOperations()
	stringOperations()
}

func integerOperations() {
	fmt.Println()
	fmt.Println("Integer Operations")
	var (
		a int = 10
		b int = 5
	)

	// Addition
	sum := a + b
	fmt.Println("Sum:", sum)

	// Subtraction
	diff := a - b
	fmt.Println("Difference:", diff)

	// Multiplication
	product := a * b
	fmt.Println("Product:", product)

	// Division
	quotient := a / b
	fmt.Println("Quotient:", quotient)

	// Remainder
	remainder := a % b
	fmt.Println("Remainder:", remainder)

	// Increment
	a++
	fmt.Println("Incremented a:", a)

	// Decrement
	b--
	fmt.Println("Decremented b:", b)
}

func floatingOperations() {
	fmt.Println()
	fmt.Println("Floating Operations")
	var (
		a float32 = 10.5
		b float32 = 5.5
	)

	// Addition
	sum := a + b
	fmt.Println("Sum:", sum)

	// Subtraction
	diff := a - b
	fmt.Println("Difference:", diff)

	// Multiplication
	product := a * b
	fmt.Println("Product:", product)

	// Division
	quotient := a / b
	fmt.Println("Quotient:", quotient)
}

func booleanOperations() {
	fmt.Println()
	fmt.Println("Boolean Operations")
	var (
		a bool = true
		b bool = false
	)

	// AND
	and := a && b
	fmt.Println("AND:", and)

	// OR
	or := a || b
	fmt.Println("OR:", or)

	// NOT
	not := !a
	fmt.Println("NOT:", not)
}

func complexOperations() {
	fmt.Println()
	fmt.Println("Complex Operations")
	var (
		a complex64 = 10 + 5i
		b complex64 = 5 + 10i
	)

	// Addition
	sum := a + b
	fmt.Println("Sum:", sum)

	// Subtraction
	diff := a - b
	fmt.Println("Difference:", diff)

	// Multiplication
	product := a * b
	fmt.Println("Product:", product)

	// Division
	quotient := a / b
	fmt.Println("Quotient:", quotient)
}

func stringOperations() {
	fmt.Println()
	fmt.Println("String Operations")
	var (
		a string = "Hello"
		b string = "World"
	)

	// Concatenation
	concatenated := a + " " + b
	fmt.Println("Concatenated:", concatenated)
}
