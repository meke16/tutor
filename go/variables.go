package main

import "fmt"

func variablesDemo() {
	//vaiable decalre in two ways in go 

	var messageWithType string = "Hello, Go Variables!"					 // Method 1: Explicit type declaration with way one variable declaration
	var messageWithTypeInference = "Hello, Go with Type Inference!"		 // Method 1: Type inference with way one variable declaration
	count := 42                                							 // Method 2: Type inference 

	/* 
		Note: It is not possible to declare a variable using :=, without assigning a value to it. i.e the following will result in a compilation error:
		variableWithoutValue :=  // This will cause a compilation error
		but with method 1 it is possible with explicit type declaration like below
		variableWithoutValue string  // This is valid
	*/

	var x int
	var name string
	fmt.Println(messageWithType)
	fmt.Println(messageWithTypeInference)
	fmt.Println("Count:", count)
	x=10
	name="Gopher"
	fmt.Println("Value of Name:", name)
	fmt.Println("Value of x:", x)
}


/* 

There are some small differences between the var var :=:

var													:=

Can be used inside and outside of functions			Can only be used inside functions
Variable declaration and value assignment can be 	done separately	Variable declaration and value assignment cannot be done separately (must be done in the same line)


*/

var a int
var b int = 2
var c = 3
//d := 3 // This will cause a compilation error
func vaiableOutSideOfFunc() {
  d := 3 // This is valid
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}