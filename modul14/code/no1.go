package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}
}

func cekJarak(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	selisih := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != selisih {
			return false
		}
	}

	return true
}

func main() {
	var x int
	var arr []int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		arr = append(arr, x)
	}

	insertionSort(arr)

	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	if cekJarak(arr) {
		fmt.Println("Data berjarak tetap")
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
