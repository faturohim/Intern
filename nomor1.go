package main

import "fmt"

func tampilkanBilanganCacah(n int) {
	count := 0
	i := 1

	for count < n {
		if i%3 == 0 || i%7 == 0 {
			fmt.Print(i)
			if i%3 == 0 && i%7 == 0 {
				fmt.Print("Z")
			}
			fmt.Println()
			count++
		}
		i++
	}
}

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)

	fmt.Println("Bilangan cacah kelipatan 3 atau 7: ")
	tampilkanBilanganCacah(N)
}
