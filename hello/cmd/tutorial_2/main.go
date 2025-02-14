package main

import "fmt"

func main() {

	var int8Num int8 = 127	
	var float32Num float32 = 12345678.9
	
	//	// MUST be same data type this will fail
	// var result float32 = float32Num - int8Num

	// Must cast int to float
	var result float32 = float32Num - float32(int8Num)
	fmt.Println(result)
	
	// int division trunc
	var int8Num1 int8 = 9
	var int8Num2 int8 = 5
	fmt.Println(int8Num1/int8Num2)
	fmt.Println(int8Num1%int8Num2)

	var myString string = "Hello World!"
	var myString2 string = `Hello 
	World!`
	fmt.Println(myString)
	fmt.Println(myString2)
}
