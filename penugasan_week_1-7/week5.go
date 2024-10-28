/*
#Soal nomor 8
Buatlah algoritma untuk menghitung nilai faktorial dari suatu bilangan bulat positif.
Masukan terdiri dari sebuah bilangan bulat positif n.
Keluaran berupa sebuah bilangan bulat yang menyatakan factorial dari n.

#pseucode
Algoritma
hasil = 1
perulangan faktorial(a)

	i = 1;i <=a;i++ {
	    hasil = hasil * i
	}

output(hasil)
endprogram
*/
package main

import (
	"fmt"
)

func main() {

	var a, hasil int

	fmt.Print("Masukan angka: ")
	fmt.Scan(&a)

	hasil = 1

	for i := 1; i <= a; i++ {
		hasil = hasil * i
	}

	fmt.Print(hasil)
}
