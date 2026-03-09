package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Masukkan suhu Celsius: ")
	fmt.Scan(&c)

	f := (c * 9 / 5) + 32
	r := c * 4 / 5
	k := c + 273

	fmt.Printf("Derajat Reamur: %.2f\n", r)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", f)
	fmt.Printf("Derajat Kelvin: %.2f\n", k)
}