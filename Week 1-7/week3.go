/*
#Soal nomor 1
Sebuah program digunakan untuk mengecek suatu bilangan adalah ganjil atau bukan.
Masukan terdiri dari bilangan bulat positif.
Keluaran terdiri dari boolean yang menyatakan true apabila bilangan adalah ganjil, atau false apabila sebaliknya
#pseucode

Algoritma
Mulai

	Input bilangan positif
	angka%2 != 0
	Output boolean

endprogram
*/
package main

import "fmt"

func main() {
	var angka int
	fmt.Printf("Masukan bilangan bulat positif:")
	fmt.Scan(&angka)

	fmt.Println(angka%2 != 0)
}
