package main

import "fmt"

func main() {
	// 1. Boolean
	fmt.Println(true)
	fmt.Println(false)

	// 2. Integers
	fmt.Println(42)
	fmt.Println(-7)
	fmt.Println(0b1010) // binary (10 in decimal)
	fmt.Println(0o12)   // octal (10 in decimal)
	fmt.Println(0xA)    // hexadecimal (10 in decimal)

	// 3. Floating-point numbers
	fmt.Println(3.14)  //  float
	fmt.Println(-0.5)  // negative float
	fmt.Println(2.0e3) // scientific notation (2000)

	// 4. Complex numbers
	fmt.Println(1 + 2i)        // literal form
	fmt.Println(complex(3, 4)) // using complex() function

	// 5. Strings
	fmt.Println("Hello, Go!") // normal string
	fmt.Println(`This is
a raw
string`) // keeps line breaks and no escaping

	// 6. Runes (Unicode characters)
	fmt.Println('A') // rune for character A
	fmt.Println('🔥') // rune for emoji
}
