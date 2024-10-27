package main

import "fmt"

func main() {
	var r, luas float64
	const phi = 3.14

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas = phi * r * r
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", r, luas)
}
