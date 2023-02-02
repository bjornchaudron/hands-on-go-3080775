// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var s any = "Hello"
	fmt.Println(s.(string))

	// perform a type assertion and handle the error
	var x int = 6
	assertInt(s)
	assertInt(x)
}

func assertInt(v any) {
	if _, ok := v.(int); !ok {
		fmt.Printf("%v is not an int\n", v)
	}
}
