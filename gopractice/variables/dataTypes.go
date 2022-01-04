package variables

import "fmt"

// DataTypes -
func DataTypes() {

	var snow bool = false
	fmt.Println("Data type snow is boolean = ", snow)

	var name string = "Go Lang"
	fmt.Println("Data type name is String = ", name)

	// var num1 int8 = 256 // overflow because exceeds range of int8
	// range of int8 is -128 to 127
	var num1 int8 = 120
	fmt.Println("Data type num1 is int8 = ", num1)

	// var num2 int16 = 50000 // overflow because exceeds range of int8
	// range of int16 is -32768 to 32767
	var num2 int16 = 32767
	fmt.Println("Data type num2 is int16 = ", num2)

	// var num3 int32 = 523435235345325324534 // overflow because exceeds range of int8
	// range of int16 is -2^31 to 2^31 - 1
	var num3 int32 = 324003242
	fmt.Println("Data type num3 is int32 = ", num3)

	// var num4 int64 = 32400324239523435235345325324534 // overflow because exceeds range of int8
	// range of int64 is -2^63 to 2^63 - 1
	var num4 int64 = 3395234352240032420
	fmt.Println("Data type num4 is int64 = ", num4)

}
