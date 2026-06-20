package main

import "fmt"

var data string
var pos int

func start() {
	pos = 0
}

func maju() {
	pos++
}

func eop() bool {
	return data[pos] == '.'
}

func cc() byte {
	return data[pos]
}

func main() {
	fmt.Print("Masukkan karakter (akhiri dengan .): ")
	fmt.Scan(&data)

	start()

	jumlahKarakter := 0
	jumlahA := 0
	jumlahLE := 0

	var prev byte

	for !eop() {

		jumlahKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if prev == 'L' && cc() == 'E' {
			jumlahLE++
		}

		prev = cc()
		maju()
	}

	frekuensiA := 0.0
	if jumlahKarakter > 0 {
		frekuensiA = float64(jumlahA) / float64(jumlahKarakter)
	}

	fmt.Println("Jumlah karakter :", jumlahKarakter)
	fmt.Println("Jumlah huruf A :", jumlahA)
	fmt.Printf("Frekuensi A : %.2f\n", frekuensiA)
	fmt.Println("Jumlah kata LE :", jumlahLE)
}