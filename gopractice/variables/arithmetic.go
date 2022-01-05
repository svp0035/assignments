package variables

import "fmt"

// ArithmeticInt -
func ArithmeticInt() {
	a := 10
	b := 20

	c := a + b                   // Addition Operator
	d := a - b                   // Subtraction Operator
	e := float64(a) / float64(b) // Division Operator
	f := a * b                   // Multiplication Operator
	g := a % b                   // Modulus Operator

	fmt.Println(c, d, e, f, g)

}

// ArithmeticFloat -
func ArithmeticFloat() {
	a := 10.5
	b := 20.9

	c := a + b                   // Addition Operator
	d := a - b                   // Subtraction Operator
	e := float64(a) / float64(b) // Division Operator
	f := a * b                   // Multiplication Operator
	// g := a % b  Modulus operator does not work with floats

	fmt.Println(c, d, e, f)

	// h := string(f)   cannot convert f (type float64) to type string
	h := int(f)
	fmt.Println(h)

}
