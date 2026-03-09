package main

import "fmt"

func main() {
	var bunga string
	pita := ""
	jumlah := 0

	for {
		fmt.Printf("Bunga %d: ", jumlah+1)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita += bunga + " - "
		jumlah++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}