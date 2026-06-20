package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n int
	fmt.Scan(&n)

	dalamPizza := 0

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			dalamPizza++
		}
	}

	fmt.Println("Topping pada Pizza:", dalamPizza)
}