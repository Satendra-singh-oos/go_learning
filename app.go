package main

import "fmt"

func hello() string {
	return "Hello World From Function"
}

func main() {
	fmt.Println("Hello World ")

	fmt.Println(hello())
}
