// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// simple greet function
func greet() string {
	return "Hello!"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// greetWithName returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = fmt.Sprintf("Hello, my name is %s and I am %s years old.", name, strconv.Itoa(age))
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func main() {
	// invoke greet function
	fmt.Println(greet())

	// invoke greetWithName function
	fmt.Println(greetWithName("Toni"))

	// invoke greetWithNameAndAge function
	fmt.Println(greetWithNameAndAge("Toni", 31))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
