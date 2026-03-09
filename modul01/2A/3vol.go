package main

import "fmt"

func main() {
	var r float64
	const pi = 3.1415926535

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * pi * r * r * r
	luas := 4 * pi * r * r

	fmt.Printf("Volume bola = %.4f\n", volume)
	fmt.Printf("Luas permukaan bola = %.4f\n", luas)
}