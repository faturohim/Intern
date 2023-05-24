package main

import (
	"fmt"
	"unicode"
)

func cekKataSandi(kataSandi string) string {

	if len(kataSandi) < 8 || len(kataSandi) > 32 {
		return "Kata sandi harus memiliki panjang antara 8 hingga 32 karakter"
	}

	if unicode.IsDigit(rune(kataSandi[0])) {
		return "Karakter awal tidak boleh angka"
	}

	hasDigit := false
	hasUpper := false
	hasLower := false

	for _, char := range kataSandi {
		switch {
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		}
	}

	if !hasDigit {
		return "Harus memiliki angka"
	}
	if !hasUpper || !hasLower {
		return "Harus memiliki huruf kapital dan huruf kecil"
	}

	return "Kata sandi valid"
}

func main() {
	var kataSandi string
	fmt.Print("Masukkan kata sandi: ")
	fmt.Scanln(&kataSandi)

	status := cekKataSandi(kataSandi)
	fmt.Println(status)
}
