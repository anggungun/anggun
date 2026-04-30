package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&N)

	var berat [1000]float64

	for i := 0; i < N; i++ {
		fmt.Print("Masukkan berat ke-", i+1, ": ")
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
