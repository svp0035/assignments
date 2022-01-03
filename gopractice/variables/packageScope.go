package variables

import "fmt"

var z int = 100 // variable is declared in the package so it can be accessible everywhere in "variables package"

func PackageVariables() {
	fmt.Println("z: = ", z)

}
