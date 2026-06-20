package main

import "fmt"

func main() {
	var x, s string
	var n int

	fmt.Scan(&x)
	fmt.Scan(&n)

	ditemukan := false
	posisi := -1
	jumlah := 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&s)

		if s == x {
			jumlah++

			if !ditemukan {
				ditemukan = true
				posisi = i
			}
		}
	}

	fmt.Println("Ada?", ditemukan)

	if ditemukan {
		fmt.Println("Posisi pertama:", posisi)
	} else {
		fmt.Println("Posisi pertama: tidak ditemukan")
	}

	fmt.Println("Jumlah kemunculan:", jumlah)

	if jumlah >= 2 {
		fmt.Println("Sedikitnya dua kali: Ya")
	} else {
		fmt.Println("Sedikitnya dua kali: Tidak")
	}
}