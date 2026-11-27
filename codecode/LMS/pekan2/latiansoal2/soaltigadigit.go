package main 

import "fmt"

func main () {
	var x int

	fmt.Scan(&x)

	d3 := x % 10
	d2 := (x / 10 ) % 10 
	d1 := x / 100
	
	fmt.Println(d1, d2, d3)
}