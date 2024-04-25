package variables

import (
	"fmt"
)

// read further at https://go.dev/blog/type-inference

func TypeInference() {
	/* Wikipedia defines type inference as follows:
	Type inference is the ability to automatically deduce,
	either partially or fully, the type of an expression at compile time.
	The compiler is often able to infer the type of a variable
	or the type signature of a function, without explicit type annotations having been given. */

	// The key phrase here is “automatically deduce … the type of an expression”.
	// Go supported a basic form of type inference from the start:
	expr := 42    // the type of expr is int
	var x1 = expr // the type of x is the type of expr
	x2 := expr

	/* No explicit types are given in these declarations,
	and therefore the types of the constant and variables
	x1, x2 and x3 on the left of = and := are the types of the
	respective initialization expressions, on the right.
	We say that the types are inferred from (the types of)
	their initialization expressions.
	With the introduction of generics in Go 1.18,
	Go’s type inference abilities were significantly expanded. */

	/* In non-generic Go code, the effect of leaving away types is
	most pronounced in a short variable declaration ( := ).
	Such a declaration combines type inference and a little bit of
	syntactic sugar—the ability to leave away the var keyword—into one very compact statement.
	Consider the following map variable declaration: */
	var x3 map[string]int = map[string]int{}
	// vs
	x4 := map[string]int{}
	// The type of x4 is map[string]int, just like the type of x3.
	// The type of the map literal on the right of = is map[string]int,
	// and the type of x4 is inferred from that.

	/* The type of a variable declared with := is always 
	inferred from the type of the initialization expression. */

	/* Omitting the type on the left of := removes repetition and at the same time increases readability. */

	fmt.Println()
	fmt.Println("Type Inference")
	fmt.Println("x1:", x1)
	fmt.Println("x2:", x2)
	fmt.Println("x3:", x3)
	fmt.Println("x4:", x4)
}
