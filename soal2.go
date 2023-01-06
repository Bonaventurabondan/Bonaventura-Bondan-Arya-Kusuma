package main

import (
	"fmt"
)

type jumlahElemen [999]string

func main() {
	var (
		arr jumlahElemen
		n   int
		str string
	)
	arr[0] = "satu"
	arr[1] = "dua"
	arr[2] = "tiga"
	arr[3] = "empat"
	arr[4] = "lima"
	n = len(arr)
	fmt.Scanln(&str)
	fmt.Println(seqSearch(arr, n, str))
}

func seqSearch(T jumlahElemen, N int, X string) int {
	var (
		ketemu int
	)

	ketemu = -1
	i := 0
	for ketemu == -1 && i < N {
		if T[i] == X {
			ketemu = i
		}
		i++

	}
	return ketemu

}
