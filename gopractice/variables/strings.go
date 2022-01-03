package variables

import "fmt"

// VarDeclarationString is declaring Strings assignment
func VarDeclarationString() {

	fmt.Println("This function will declare STRING variables")

	// short hand declaration where initializatoin and declaration is done in same line
	a, b := "Hello There", "GoLang"
	fmt.Println(a, b)

	// initalise first then declare a value
	var c, d string
	c, d = "Sage", "IT"
	fmt.Println(c, d)

	// initalize and declare variables
	var e, f string = "New", "Year"
	fmt.Println(e, f)

}
