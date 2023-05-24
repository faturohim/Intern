package main

import (
	"fmt"
	"strings"
)

func cariKata(teks string) {
	teks = strings.ToLower(teks)
	kata := strings.Split(teks, " ")

	for i := 0; i < len(kata); i++ {
		if kata[i] == "sang" && kata[i+1] == "gajah" {
			fmt.Print("sang gajah - ")
		} else if kata[i] == "serigala" {
			fmt.Print("serigala - ")
		} else if kata[i] == "harimau" {
			fmt.Print("harimau - ")
		}
	}

	fmt.Println()
}

func main() {
	teks := "Berikut adalah kisah sang gajah. Sang gajah memiliki teman serigala bernama Doe. Sang Gajah sering dibela oleh serigala ketika harimau mendekati gajah."

	cariKata(teks)
}
