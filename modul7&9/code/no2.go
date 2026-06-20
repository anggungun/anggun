package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	var arr [100]int

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\nIndeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\nIndeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	var x int
	fmt.Print("\nKelipatan indeks: ")
	fmt.Scan(&x)

	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var hapus int
	fmt.Print("\nIndeks yang dihapus: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	var jumlah int
	for i := 0; i < n; i++ {
		jumlah += arr[i]
	}

	rata := float64(jumlah) / float64(n)
	fmt.Printf("\nRata-rata: %.2f\n", rata)

	var total float64
	for i := 0; i < n; i++ {
		total += math.Pow(float64(arr[i])-rata, 2)
	}

	sd := math.Sqrt(total / float64(n))
	fmt.Printf("Standar Deviasi: %.2f\n", sd)
}