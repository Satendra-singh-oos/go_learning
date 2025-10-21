package main

import "fmt"

// in go we only have for
// for-> only construct in go for looping

func main() {
	// While-style loop (condition only)
	i := 1
	for i <= 3 {
		fmt.Println("While loop:", i)
		i++
	}

	//  Classic for loop (init; condition; post)
	for j := 1; j <= 3; j++ {
		fmt.Println("Classic for loop:", j)
	}

	// Infinite loop (break when done)
	k := 1
	for {
		fmt.Println("Infinite loop:", k)
		k++
		if k > 3 {
			break
		}
	}

	// 4️ Range loop (over slices, arrays, maps, strings)
	names := []string{"Go", "Python", "JavaScript"}
	for index, value := range names {
		fmt.Printf("Range loop: index=%d, value=%s\n", index, value)
	}

	//  Range over string (runes)
	word := "Go🔥"
	for idx, runeVal := range word {
		fmt.Printf("Rune loop: idx=%d, char=%c\n", idx, runeVal)
	}
}
