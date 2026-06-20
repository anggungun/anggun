package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	Sisi1 int
	Sisi2 int
	Nilai int
	Balak bool
}

type Dominoes struct {
	Kartu  [28]Domino
	Jumlah int
}

// Membuat 1 set kartu domino
func buatDominoes() Dominoes {
	var d Dominoes
	idx := 0

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.Kartu[idx] = Domino{
				Sisi1: i,
				Sisi2: j,
				Nilai: i + j,
				Balak: i == j,
			}
			idx++
		}
	}

	d.Jumlah = 28
	return d
}

// Soal 1b
func kocokKartu(d *Dominoes) {
	for i := d.Jumlah - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Kartu[i], d.Kartu[j] = d.Kartu[j], d.Kartu[i]
	}
}

// Soal 1c
func ambilKartu(d *Dominoes) Domino {
	if d.Jumlah == 0 {
		return Domino{}
	}

	d.Jumlah--
	return d.Kartu[d.Jumlah]
}

// Soal 1d
func gambarKartu(k Domino, suit int) int {
	if suit == 1 {
		return k.Sisi1
	}
	return k.Sisi2
}

// Soal 1e
func nilaiKartu(k Domino) int {
	return k.Nilai
}

// Soal 2a
func galiKartu(d *Dominoes, target Domino) {
	fmt.Println("\nMencari kartu yang memiliki suit sama dengan:")
	fmt.Printf("(%d,%d)\n", target.Sisi1, target.Sisi2)

	for d.Jumlah > 0 {
		k := ambilKartu(d)

		if k.Sisi1 == target.Sisi1 ||
			k.Sisi1 == target.Sisi2 ||
			k.Sisi2 == target.Sisi1 ||
			k.Sisi2 == target.Sisi2 {

			fmt.Printf("Ditemukan kartu (%d,%d)\n", k.Sisi1, k.Sisi2)
			return
		}
	}

	fmt.Println("Tidak ditemukan kartu yang cocok")
}

// Soal 2b
func sepasangKartu(k1, k2 Domino) bool {
	return nilaiKartu(k1)+nilaiKartu(k2) == 12
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dominoes := buatDominoes()

	fmt.Println("Jumlah kartu awal:", dominoes.Jumlah)

	kocokKartu(&dominoes)

	k1 := ambilKartu(&dominoes)
	k2 := ambilKartu(&dominoes)

	fmt.Printf("Kartu 1 : (%d,%d)\n", k1.Sisi1, k1.Sisi2)
	fmt.Printf("Kartu 2 : (%d,%d)\n", k2.Sisi1, k2.Sisi2)

	fmt.Println("Nilai kartu 1 :", nilaiKartu(k1))
	fmt.Println("Nilai kartu 2 :", nilaiKartu(k2))

	if k1.Balak {
		fmt.Println("Kartu 1 adalah balak")
	}

	if sepasangKartu(k1, k2) {
		fmt.Println("Total nilai kedua kartu = 12")
	} else {
		fmt.Println("Total nilai kedua kartu ≠ 12")
	}

	galiKartu(&dominoes, k1)
}