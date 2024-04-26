package general

import "fmt"

/*
Iota is an identifier which is used with constant and which can
simplify constant definitions that use auto increment numbers.
The IOTA keyword represent integer constant starting from zero.
So essentially it can be used to create effective constant in Go.
They can also be used to create enumeration like effect in Go.
*/

/*
Suppose you want to define a set of constants that represent the
days of the week. You can define the first constant as 0 and then
keep incrementing it by 1 for each of the next constants.
*/

// Define the days of the week without using iota
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// The same can be done using iota
// Define the days of the week using iota
const (
	Sunday1 = iota
	Monday1
	Tuesday1
	Wednesday1
	Thursday1
	Friday1
	Saturday1
)

/*
So iota is:
1. A counter which starts with zero
2. Increases by 1 after each line
3. Is only used with constant
*/

// some other ways to use iota

// this works the same way as the days of the week
const (
	x1 = iota
	x2 = iota
	x3 = iota
)

// iota keyword can be skipped in the definition and it will still work
const (
	x4 = iota
	x5
	x6 = iota
)

// iota keyword cannont be skipped in the definition of the first constant
// this will not work
// cosnt (
// 	x7
// 	x8 = iota
// 	x9
// )

// adding spaces and comments does not affect the iota counter
const (
	x7 = iota

	x8
	// comment
	x9
)

// iota increment can be skipped using the blank identifier
const (
	x10 = iota
	_
	x11
	x12
)

// we can also use expressions to manipulate the iota counter
const (
	x13 = iota
	x14 = iota + 5 // 1 + 5
	x15 = iota + 6 // 2 + 6
)
const (
	x16 = iota + 100
	x17
	x18
)

func TestIota() {
	// Print the days of the week
	fmt.Println()
	println("Without Iota")
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	println("With Iota")
	println(Sunday1, Monday1, Tuesday1, Wednesday1, Thursday1, Friday1, Saturday1)

	fmt.Println()
	fmt.Println(x1, x2, x3)
	fmt.Println(x4, x5, x6)
	fmt.Println(x7, x8, x9)
	fmt.Println(x10, x11, x12)
	fmt.Println(x13, x14, x15)
	fmt.Println(x16, x17, x18)
}
