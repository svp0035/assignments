package variables

import "fmt"

// VarDeclarationInt function is for variables assignment
func VarDeclarationInt() {

	fmt.Println("This function will declare INT variables")

	// short hand declaration where initializatoin and declaration is done in same line

	a, b := 10, 11
	fmt.Println(a, b)

	// initalise first then declare a value
	var c, d int
	c, d = 12, 13
	fmt.Println(c, d)

	// initalize and declare variables
	var e, f int = 14, 15
	fmt.Println(e, f)
}
