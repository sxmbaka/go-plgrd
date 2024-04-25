package variables

import "fmt"

func TestAllVariables() {
	integerTypes()
	octalAndHexadecimalTypes()
	floatingTypes()
	booleanType()
	complexTypes()
	stringType()
}

func integerTypes() {
	// int8   ranges from    -128 to 127
	// uint   ranges from       0 to 255
	// int16  ranges from   -2^15 to 2^15 -1
	// uint16 ranges from       0 to 2^16 -1
	// int32  ranges from   -2^31 to 2^31 -1
	// uint32 ranges from       0 to 2^32 -1
	// int64  ranges from   -2^63 to 2^63 -1
	// uint64 ranges from       0 to 2^64 -1

	// byte = uint8
	// byte   ranges from       0 to 255

	// rune = int32
	// rune   ranges from 	-2^31 to 2^31 -1

	// int and uint are platform dependent
	// int maybe int32 or int64
	// uint maybe uint32 or uint64

	var (
		x1  int8   = -8
		x2  uint8  = 8
		x3  byte   = 8
		x4  int16  = -16
		x5  uint16 = 16
		x6  int32  = -32
		x7  rune   = -32
		x8  uint32 = 32
		x9  int64  = -64
		x10 uint64 = 64
		x11 int    = -32064
		x12 uint   = 32064
	)
	fmt.Println()
	fmt.Println("x1(int8)           :", x1)
	fmt.Println("x2(uint8)          :", x2)
	fmt.Println("x3(byte)(uint8)    :", x3)
	fmt.Println("x4(int16)          :", x4)
	fmt.Println("x5(uint16)         :", x5)
	fmt.Println("x6(int32)          :", x6)
	fmt.Println("x7(rune)(int32)    :", x7)
	fmt.Println("x8(uint32)         :", x8)
	fmt.Println("x9(int64)          :", x9)
	fmt.Println("x10(uint64)        :", x10)
	fmt.Println("x11(int)           :", x11)
	fmt.Println("x12(uint)          :", x12)
}

func floatingTypes() {
	// float32 ranges from -3.4e+38 to 3.4e+38
	// float64 ranges from -1.7e+308 to +1.7e+308

	var (
		x1 float32 = 32.32
		x2 float64 = 64.64
	)

	fmt.Println()
	fmt.Println("x1(float32)        :", x1)
	fmt.Println("x2(float64)        :", x2)
}

func booleanType() {
	// bool type is used to store binary values mapped as true and false

	// different types you can declare booleans
	// this also cleverly covers all the logical operators in go
	var (
		x1 bool = true
		x2      = false
		x3      = 3 == 4
		x4      = 3 < 4
		x5      = 3 > 4
		x6      = 3 >= 4
		x7      = 3 <= 4
		x8      = 3 != 4
		x9      = !false
	)

	fmt.Println()
	fmt.Println(x1, x2, x3, x4, x5, x6, x7, x8, x9)
}

func octalAndHexadecimalTypes() {
	// in go, you can declare octal numbers using prefix 0
	// and hexadecimal numbers using the prefix 0x or 0X
	// following is a complete example of integer types

	var (
		// this will take the type int
		x1 = 34
		// but if we add a 0 before the value it is treated as an octal value
		x2 = 034
		// and if we add a 0x or 0X before the value it is treated as a hexadecimal value
		x3 = 0xA3F
		x4 = 0xAf3D
	)

	fmt.Println()
	fmt.Printf("int typed          : %d\n", x1)
	fmt.Printf("octal              : %#o\n", x2)
	fmt.Printf("hex (0x)           : %#x\n", x3)
	fmt.Printf("hex (0X)           : %#x\n", x4)
}

func complexTypes() {
	// complex64 has float32 real and imaginary parts
	// complex128 has float64 real and imaginary parts

	var (
		x1 complex64  = 3.2 + 12i
		x2 complex128 = 3.2 + 12i

		a float32 = 3.2
		b float32 = 12
		c float64 = 3.2
		d float64 = 12

		// you can also declare complex numbers using the complex function
		// complex function takes two arguments, the real and imaginary parts
		// and returns a complex number
		x3 complex64  = complex(a, b)
		x4 complex128 = complex(c, d)

		// this does not work if you try to declare a complex number using variables of different float types
		// you have to convert the variables to the same type (if possible) before using them
		// e.g., complex(a, d) or complex(c, b) give errors
	)

	fmt.Println()
	fmt.Println("x1(complex64)      :", x1)
	fmt.Println("x2(complex128)     :", x2)
	fmt.Println("x3(complex64)      :", x3)
	fmt.Println("x4(complex128)     :", x4)
}

func stringType() {
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var x1 = "Steve Jobs"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var x2 = `Steve Jobs was an American entrepreneur and inventor.
		     He was the CEO and co-founder of Apple Inc.`
	// x2 will contain the string as it is, with the newline and tab characters

	fmt.Println()
	fmt.Println("x1(string)         :", x1)
	fmt.Println("x2(raw string)     :", x2)
}
