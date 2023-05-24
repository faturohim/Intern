package main

import (
	"fmt"
)

func printPattern(n int) {
	pattern := make([][]rune, n)

	for i := 0; i < n; i++ {
		pattern[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			pattern[i][j] = 'X'
		}
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			pattern[i][j] = 'O'
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%c", pattern[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N (bilangan ganjil): ")
	fmt.Scan(&n)

	if n%2 != 0 && n > 0 {
		fmt.Println("Pola: ")
		printPattern(n)
	} else {
		fmt.Println("N harus merupakan bilangan ganjil dan lebih dari 0.")
	}
}
