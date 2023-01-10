// types/variables/begin/main.go
package main

import "fmt"

// declare package-level variables of type int
var a, b, c int

// declare package-level variables of type bool and override the default values (also known as "zero")
var x, y, z = true, false, true

func main() {
	// print ints
	fmt.Println("a, b, c:", a, b, c)

	// override the default value of a package-level variable
	c = 1_000_000
	fmt.Printf("d: %d\n", c)

	// declare and initialize variables of similar names as booleans but of local scope
	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, z:", x, y, z)

	// print the package-level variables x, y, and z
	printVars()
}

func printVars() {
	fmt.Println("x, y, z:", x, y, z)
}