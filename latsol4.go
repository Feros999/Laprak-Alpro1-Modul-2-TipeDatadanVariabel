package main

import "fmt"

func main() {
	var fahrenheit, celsius float64
	
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	celsius = (fahrenheit - 32) * 5 / 9
	fmt.Printf("Suhu dalam Celcius: %.2f\n", celsius)
}
