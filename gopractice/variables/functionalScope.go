package variables

import "fmt"

// FunctionalScope
func FunctionalScope() {
	var a int = 10
	fmt.Println("a:", a)
	fmt.Println("z package variable is: ", z)                         // z is a package variable so can be executed here even though it is defined in another file
	fmt.Println("global variable in functoin is accessed : ", Global) // Global variables is accessed
}

// FunctionalScope2
func FunctionalScope2() {
	//fmt.Println(a) Code will not execute because "a" is declared in a function and is not accessible outside
}
