package main

import "fmt"

/*
-  If you declare a variable but don’t use it, the compiler will throw an error.
- 3 ways to define variables in Go:
 1. var with type
 2. var without type (type inference)
 3. short declaration :=

-  Constants use 'const' keyword
*/

const outSide = "Golang From Outside Of Function"

func main() {
	//  var with type
	var name string
	name = "Hello Go!"
	fmt.Println(name)

	// var without type (type inference)
	var newName = "Hello Go Again"
	fmt.Println(newName)

	//  Boolean type
	var isBool = true
	fmt.Println(isBool)

	// Integer type
	var age int = 25
	fmt.Println(age)

	// Short declaration (only inside functions)
	country := "India"
	fmt.Println(country)

	// Constant (cannot change value)
	const PI = 3.14
	fmt.Println(PI)

	// float
	// var price float32 = 50.5
	//var price =50.5
	price := 50.5
	fmt.Println("Initial Value", price)
	price = 90
	fmt.Println("Update Value", price)

	println(outSide)

	// group multiple const
	const (
		port = 5000
		host = "localhost:"
	)

	fmt.Println(host, port)
}
