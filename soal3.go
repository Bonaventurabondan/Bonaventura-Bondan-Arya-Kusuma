package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])

	}
}
func posisi(n, k int) int {
	var i int = 0
	for i < n && data[i] != k {
		i++
	}
	if i < n {
		return i
	} else {
		return -1
	}

}

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	isiArray(n)

	hasil := posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
