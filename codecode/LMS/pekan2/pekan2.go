package main

import (
	"fmt"
	"math"
)

func main () {
	var a, b, jumlah int

	fmt.Scan(&a, &b)
	jumlah = a + b
	fmt.Println(jumlah)
	fmt.Println("hello") //dari fmt/format
	fmt.Println(math.Pow(5.0, 2.0)) //dari math/mathematic(5 pangkat 2)
	universal () //bagian ini untuk memanggil bagian func universal, karena func akan selalu utama
}

//blok ini dipanggil dari func main
func universal () {
	var nama string = "rolland"
	fmt.Println("nama saya:", nama)
}