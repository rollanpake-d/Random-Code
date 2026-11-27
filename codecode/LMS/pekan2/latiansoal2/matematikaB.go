package main

import "fmt"

func main () {
	var x float64

	fmt.Scan(&x)

	f := (x * x  * x + 3.0 * x) / (x * x * x * x - 3.0 * x * x + 4.0)

	fmt.Println(f)
}