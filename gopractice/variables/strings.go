package variables

import "fmt"

// StringDeclaration is declaring Strings assignment
func VarDeclarationString() {

	fmt.Println("This function will declare STRING variables")

	a, b := "Hello There", "GoLang"
	fmt.Println(a, b)

	var c, d string
	c, d = "Sage", "IT"
	fmt.Println(c, d)

	var e, f string = "New", "Year"
	fmt.Println(e, f)

}
