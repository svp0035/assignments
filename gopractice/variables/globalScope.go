package variables

import "fmt"

//Global Exported  variable
var Global int = 50 // capitalized Global variable so it can be accessed everywhere

//GlobalVariables -
func GlobalVariables() {

	fmt.Println("Global variable: ", Global)
}
