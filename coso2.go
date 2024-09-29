package main

import "fmt"

func main() {
	var x, fx float64
	// masukkan nilai x
	fmt.Println("input x")
	fmt.Scan(&x)
	// fx = 2 / (x+5) + 5
	fx = 2/(x+5) + 5
	fmt.Println(fx)
}
