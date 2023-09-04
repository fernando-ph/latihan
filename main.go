package main

import (
	"fmt"
)

func main() {
	//ngebuat variabel dengan var
	var Buat = "sayang"
	fmt.Println(Buat)
	// ngebuat variabel dengan angka
	var buat = 23
	fmt.Println(buat)
	//deklarasi var dengan tipe data
	var age int
	age = 345
	fmt.Println(age)
	//deklarasi var dengan titik dua
	//titik dua itu bisa dideklarasi kan tanpa tipe data dengan contoh dibawah ini
	//tapi tidak bisa dideklarasikan 2 kali dan juga menganti tipe datanya dengan lain cth dibawah ini
	//nando := "ss"
	//nando :="ss2"
	//ini tidk bisa
	//dan satu lagi tidak bisa ganti tipe data
	//nando := 2
	//nando :="s"
	//tidak bisa tidak bisa ganti tipe
	mobile := "mobile"
	fmt.Println(mobile)
	mobile = "atu"
	fmt.Println(mobile)
	mobile2 := 23
	fmt.Println(mobile2)
	//bisa mendeklarasikan variable dengan memasukan nya ke dalam kurung
	var (
		namaDepan     = "s"
		namaBelakakng = "b"
		kelas         = 12
		bool          = true
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakakng)
	fmt.Println(kelas)
	fmt.Println(bool)

}
