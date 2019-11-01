package main

import (
	"fmt"
)

func main() {
	var username, password string
	const Uname string = "10119152"
	const Pass string = "13012002"

	salah := 0
	for salah < 3 {
		//memasukkan username dan password
		fmt.Println("<< L O G I N >>")
		fmt.Print("NIM      : ")
		fmt.Scan(&username)
		fmt.Print("Password : ")
		fmt.Scan(&password)

		//validasi ussername dan password
		if (username == Uname) && (password == Pass) {
			fmt.Println("username dan password benar, anda akan login...")
			salah = 4
		} else {
			fmt.Println("username atau password salah, Ulangi!")
			salah++
		}
	}
	if salah == 3 {
		fmt.Println("Salah lebih dari 3 kali, anda diblokir")
	}
}
