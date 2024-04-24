package variables

import (
	"fmt"
)

func IntegerOperations() {
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
