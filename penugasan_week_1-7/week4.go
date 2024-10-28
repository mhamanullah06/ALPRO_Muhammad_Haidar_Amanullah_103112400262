/*
#Soal nomor 2
Sebuah program digunakan untuk menyimpan data persegi panjang yang berisi panjang, lebar, luas, dan keliling persegi panjang.
Masukan: terdiri dari dua bilangan riil yang menyatakan panjang dan lebar dari persegi panjang.
Keluaran: terdiri dari dua bilangan riil yang menyatakan luas dan keliling dari persegi panjang.
Catatan: Gunakan tipe bentukan untuk menyimpan data persegi panjang tersebut.

#pseucode
Algoritma
Mulai
	Input p
	Input l
	luas p * l
	keliling 2* (p + l)
	Output nilai panjang
	Output nilai lebar
	Output nilai luas
	Output nilai keliling
endprogram
*/

package main

import "fmt"

func main() {
	var p, l float64

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&l)

	luas := p * l
	keliling := 2 * (p + l)

	fmt.Printf("\nInformasi Persegi Panjang:\n")
	fmt.Printf("Panjang: %.2f\n", p)
	fmt.Printf("Lebar: %.2f\n", l)
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}
