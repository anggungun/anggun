package main

import "fmt"

func main() {
	var bil, jumlah float64
	var banyak int

	fmt.Scan(&bil)

	for bil != 9999 {
		jumlah += bil
		banyak++
		fmt.Scan(&bil)
	}

	if banyak > 0 {
		fmt.Printf("Rerata = %.2f\n", jumlah/float64(banyak))
	} else {
		fmt.Println("Tidak ada data")
	}
}