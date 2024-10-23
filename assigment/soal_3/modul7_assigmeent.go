package main

import "fmt"

func main() {
	var qirat int

	fmt.Print("masukan jumlah qirat :")
	fmt.Scan(&qirat)

	fals := qirat / 6
	sisaqirat := fals % 6

	dihram := fals / 10
	sisafals := dihram % 10

	dinar := dihram / 10
	sisadihram := dinar % 10

	fmt.Printf("hasil konversi: %d dinar, %d dirham, %d fals, %d qirat\n", dinar, sisaqirat, sisafals, sisadihram)

}
