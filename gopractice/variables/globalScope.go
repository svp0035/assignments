package variables

import "fmt"

var Global int = 50 // capitalized Global variable so it can be accessed everywhere

func GlobalVariables() {

	fmt.Println("Global variable: ", Global)
}
