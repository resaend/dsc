package main

import (
	"fmt"
)

func main() {
	var pilihan int
	var angka1, angka2 float32

	//menampilkan menu pilihan kalkulator sederhana
	for pilihan != 5 {
		fmt.Println("Kalkulator sederhana")
		fmt.Println("1. Tambah")
		fmt.Println("2. Kurang")
		fmt.Println("3. Kali")
		fmt.Println("4. Bagi")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)

		//validasi pilihan
		for pilihan < 1 || pilihan > 5 {
			fmt.Println("Pilihan tidak tersedia, masukkan ulang!")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihan)
		}
		//memilih pilihan
		switch pilihan {
		case 1:
			fmt.Println("===========")
			fmt.Println("Penjumlahan")
			fmt.Println("===========")
			fmt.Print("angka pertama : ")
			fmt.Scan(&angka1)
			fmt.Print("angka kedua   : ")
			fmt.Scan(&angka2)
			fmt.Print("hasil dari ")
			fmt.Print(angka1)
			fmt.Print(" + ")
			fmt.Print(angka2)
			fmt.Print(" adalah ")
			fmt.Println(angka1 + angka2)
			fmt.Println("")
			fmt.Println("===========================")
		case 2:
			fmt.Println("===========")
			fmt.Println("Pengurangan")
			fmt.Println("===========")
			fmt.Print("angka pertama : ")
			fmt.Scan(&angka1)
			fmt.Print("angka kedua   : ")
			fmt.Scan(&angka2)
			fmt.Print("hasil dari ")
			fmt.Print(angka1)
			fmt.Print(" - ")
			fmt.Print(angka2)
			fmt.Print(" adalah ")
			fmt.Print(angka1 - angka2)
			fmt.Println("")
			fmt.Println("===========================")
		case 3:
			fmt.Println("=========")
			fmt.Println("Perkalian")
			fmt.Println("=========")
			fmt.Print("angka pertama : ")
			fmt.Scan(&angka1)
			fmt.Print("angka kedua   : ")
			fmt.Scan(&angka2)
			fmt.Print("hasil dari ")
			fmt.Print(angka1)
			fmt.Print(" x ")
			fmt.Print(angka2)
			fmt.Print(" adalah ")
			fmt.Print(angka1 * angka2)
			fmt.Println("")
			fmt.Println("===========================")
		case 4:
			fmt.Println("=========")
			fmt.Println("Pembagian")
			fmt.Println("=========")
			fmt.Print("angka pertama : ")
			fmt.Scan(&angka1)
			fmt.Print("angka kedua   : ")
			fmt.Scan(&angka2)
			fmt.Print("hasil dari ")
			fmt.Print(angka1)
			fmt.Print(" : ")
			fmt.Print(angka2)
			fmt.Print(" adalah ")
			fmt.Print(angka1 / angka2)
			fmt.Println("")
			fmt.Println("===========================")
		case 5:
		}
	}
}
