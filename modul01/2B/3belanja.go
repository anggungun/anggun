package main

import "fmt"

func main() {
	var kiri, kanan float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 || kiri+kanan > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := kiri - kanan
		if selisih < 0 {
			selisih = -selisih
		}

		fmt.Println("Sepeda motor Pak Andi akan oleng:", selisih >= 9)
	}
}