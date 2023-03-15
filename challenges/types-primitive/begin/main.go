// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	var val int = 42

	// print the value of the local variable "val"
	fmt.Printf("local val = %v\t type = %T\n", val, val)

	// print the value of the package-level variable "val"
	printVal()

	// update the package-level variable "val" and print it again
	updateVal("updated global")
	printVal()

	// print the pointer address of the local variable "val"
	fmt.Printf("local &val = %v\t type = %T\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100
	fmt.Printf("local val = %v\t type = %T\n", val, val)
}

func printVal() {
	fmt.Printf("global val = %v\t type = %T\n", val, val)
}

func updateVal(newVal string){
	val = newVal
}
