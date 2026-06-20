package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlah float64
	var piLama, piBaru float64
	i := 0

	for {
		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			jumlah += suku
		} else {
			jumlah -= suku
		}

		piBaru = 4 * jumlah

		if i > 0 {
			if math.Abs(piBaru-piLama) <= 0.00001 {
				fmt.Printf("Hasil PI sebelumnya : %.10f\n", piLama)
				fmt.Printf("Hasil PI sekarang   : %.10f\n", piBaru)
				fmt.Printf("Pada i ke : %d\n", i+1)
				break
			}
		}

		piLama = piBaru
		i++
	}
}