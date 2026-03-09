package main

import "fmt"

func main() {
	var K int
	hasil := 1.0

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	for k := 0; k <= K; k++ {
		atas := float64((4*k + 2) * (4*k + 2))
		bawah := float64((4*k + 1) * (4*k + 3))
		hasil *= atas / bawah
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}