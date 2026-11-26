package main

import "fmt"

func main() {
	var nama, nim, jurusan string

	fmt.Printf("Beri tahu nama kamu: ")
	fmt.Scan(&nama)
	
	fmt.Printf("Beri tau berapa NIM kamu: ")
	fmt.Scan(&nim)

	fmt.Printf("Beri tahu kamu dari jurusan apa: ")
	fmt.Scan(&jurusan)

	fmt.Println("Owh ternyata nama kamu adalah" ,nama, "dengan nim", nim, "dari jurusan", jurusan)

	fmt.Println("Terimakasih sudah memberikan informasi ini")
}