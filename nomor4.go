package main

import (
	"fmt"
	"sort"
)

func cariBilanganCacahTerkecil(data []int) int {

	sort.Ints(data)

	bilangan := 1
	for _, num := range data {
		if num == bilangan {
			bilangan++
		} else if num > bilangan {
			break
		}
	}

	return bilangan
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scan(&n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan bilangan ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	fmt.Println("Data:", data)
	fmt.Println("Bilangan cacah terkecil yang tidak ada adalah:", cariBilanganCacahTerkecil(data))
}
