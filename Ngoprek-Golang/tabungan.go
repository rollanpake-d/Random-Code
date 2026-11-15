package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func masukan(msg string) string {
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func daftarMenu() {
	fmt.Println("1. Tampilkan daftar tabungan")
	fmt.Println("2. Buat tabungan baru")
	fmt.Println("3. Isi tabungan")
	fmt.Println("4. Ambil isi tabungan")
	fmt.Println("5. Edit tabungan")
	fmt.Println("6. Hapus tabungan")
	fmt.Println("0. Keluar")
} 

type tabungan struct{
	nama string
	target float64
	terkumpul float64
	status string
}

func perbaruiStatus(t *tabungan) {
	if t.terkumpul >= t.target {
		t.status = "success"
	} else {
		t.status = "proses"
	}
}

func buatTabunganBaru(daftar []tabungan, namaTabungan string, targetTabungan float64) []tabungan {
	tabunganBaru := tabungan{nama: namaTabungan, target: targetTabungan, terkumpul:0.0, status: "proses"}
	fmt.Println("\nTabungan baru berhasil dibuat!\n")
	return append(daftar, tabunganBaru)
}

func tampilkanTabungan(daftar []tabungan){
	if len(daftar) == 0 {
		fmt.Printf("\nTabungan belum pernah dibuat\n\n")
	} else {
		fmt.Printf("%-4s %-25s %-18s %-18s %-18s %-10s\n", "No", "Nama", "Target", "Terkumpul", "Kekurangan", "Status")
		fmt.Println(strings.Repeat("-", 100))
		for i, v := range daftar {

			kurang := v.target - v.terkumpul
			if kurang < 0 {
				kurang = 0
			}
			fmt.Printf("%-4d %-25s %-18.0f %-18.0f %-18.0f %-10s\n", i+1, v.nama, v.target, v.terkumpul, kurang, v.status)
		}
		fmt.Println()
	}
}

func isiTabungan(daftar []tabungan, namaTabungan string){
	for i,v := range daftar{
		if(strings.EqualFold(v.nama, namaTabungan)){
			nominal, err := strconv.ParseFloat(masukan("Masukkan nominal pengisian : Rp "), 64)
			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}
			daftar[i].terkumpul += nominal
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan %q berhasil diisi!\n", v.nama)
			if daftar[i].status == "success" {
				fmt.Println("Selamat, target tabungan sudah tercapai!")
			}
			fmt.Println()
			return
		}
	}
	fmt.Println("\nNama tabungan tidak ditemukan!\n")
}

func ambilIsiTabungan(daftar []tabungan, namaTabungan string){
	for i,v := range daftar{
		if(strings.EqualFold(v.nama, namaTabungan)){
			nominal, err := strconv.ParseFloat(masukan("Nominal pengambilan : "), 64)
			if err != nil || nominal <= 0 {
				fmt.Println("Nominal harus berupa angka!")
				return
			}
			if(nominal > v.terkumpul){
				fmt.Println("Saldo tabungan tidak cukup untuk diambil")
				return
			}

			daftar[i].terkumpul -= nominal
			fmt.Printf("\nTabungan %q berhasil diambil!\n", v.nama)
			return
		}
	}
	fmt.Println("\nNama tabungan tidak ditemukan!\n")
}

func editTabungan(daftar []tabungan, namaTabungan string){
	for i := range daftar{
		if(strings.EqualFold(daftar[i].nama, namaTabungan)){
			namaBaru := masukan("Masukan nama tabungan baru : ")
			if namaBaru == "" {
				fmt.Print("\nNama tabungan tidak boleh kosong!\n\n")
			return
			}
			targetBaru, err := strconv.ParseFloat(masukan("Masukkan target baru : Rp "), 64)
			if err != nil || targetBaru <= 0 {
				fmt.Print("\nTarget harus berupa angka lebih dari 0!\n\n")
			return
			}
			daftar[i].nama = namaBaru
			daftar[i].target = targetBaru
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan berhasil diedit\n\n")
			return
		}
	}
	fmt.Println("\nNama tabungan tidak ditemukan!\n")
}

func hapusTabungan(daftar []tabungan, namaTabungan string) []tabungan {
	for i,v := range daftar{
		if(strings.EqualFold(v.nama, namaTabungan)){
			fmt.Printf("\nTabungan %q berhasil dihapus!\n\n", v.nama)
			return append(daftar[:i], daftar[i+1:]...)
		}
	}
	fmt.Printf("\nNama tabungan tidak ditemukan!\n")
	return daftar
}

func main() {
	daftarTabungan := []tabungan{}

	fmt.Println("Tabungan SiRajin")
	fmt.Println()

	for{
		daftarMenu()

		pilihan, _ := strconv.Atoi(masukan("Masukan nomor perintah :"))

		if pilihan == 0{
			break
		}

		switch pilihan {
		case 1:
			tampilkanTabungan(daftarTabungan) 
		case 2:
			nama := masukan("Masukkan nama tabungan baru : ")
			target, err := strconv.ParseFloat(masukan("Masukkan target tabungan baru : Rp "), 64)

			if err != nil {
				fmt.Printf("\nTarget tabungan harus berupa angka!\n")
			} else {
				daftarTabungan = buatTabunganBaru(daftarTabungan, nama, target)
			}
		case 3:
			nama := masukan("Masukkan nama tabungan yang diisi : ")
			isiTabungan(daftarTabungan, nama)
		case 4:
			nama := masukan("Masukkan nama tabungan yang ingin diambil : ")
			ambilIsiTabungan(daftarTabungan, nama)
		case 5:
			nama := masukan("Masukkan nama tabungan yang ingin diedit : ")
			editTabungan(daftarTabungan, nama)
		case 6:
			nama := masukan("Masukkan nama tabungan yang ingin dihapus : ")
			daftarTabungan = hapusTabungan(daftarTabungan, nama)
		}
	}
}