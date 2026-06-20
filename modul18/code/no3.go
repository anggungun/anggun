package main

import "fmt"

type Domino struct {
	Sisi1 int
	Sisi2 int
}

func bisaDisambung(k1, k2 Domino) bool {
	return k1.Sisi1 == k2.Sisi1 ||
		k1.Sisi1 == k2.Sisi2 ||
		k1.Sisi2 == k2.Sisi1 ||
		k1.Sisi2 == k2.Sisi2
}

func main() {
	kartu1 := Domino{2, 5}
	kartu2 := Domino{5, 6}

	fmt.Printf("Kartu 1 : (%d,%d)\n", kartu1.Sisi1, kartu1.Sisi2)
	fmt.Printf("Kartu 2 : (%d,%d)\n", kartu2.Sisi1, kartu2.Sisi2)

	if bisaDisambung(kartu1, kartu2) {
		fmt.Println("Kedua kartu dapat disambungkan")
	} else {
		fmt.Println("Kedua kartu tidak dapat disambungkan")
	}
}