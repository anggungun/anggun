package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan kapasitas wadah: ")
	fmt.Scan(&y)

	var ikan [1000]float64

	for i := 0; i < x; i++ {
		fmt.Print("Berat ikan ke-", i+1, ": ")
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	var wadah [1000]float64
	index := 0

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0
		count := 0

		for count < y && index < x {
			total += ikan[index]
			index++
			count++
		}

		wadah[i] = total
		fmt.Println("Total wadah", i+1, ":", wadah[i])
	}

	totalSemua := 0.0
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += wadah[i]
	}

	rata := totalSemua / float64(jumlahWadah)
	fmt.Println("Rata-rata per wadah:", rata)
}
