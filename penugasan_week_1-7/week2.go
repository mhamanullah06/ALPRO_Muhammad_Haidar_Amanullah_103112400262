/*
#Soal nomor 1
Apabila diketahui panjang dan lebar dari persegi panjang adalah p dan l maka buatlah sebuah algoritma
yang digunakan untuk menghitung luas (p*l) dan keliling (2p+2l) suatu persegi panjang

#pseucode

Algoritma

Mulai
    Input p
    Input l
    luas = p * l
    keliling = 2 * p + 2 * l
    Output luas
    Output keliling
endprogram
*/

package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("Masukan panjang:")
	fmt.Scan(&p)
	fmt.Print("Masukan lebbar:")
	fmt.Scan(&l)

	keliling := 2*p + 2*l
	luas := p * l

	fmt.Printf("Keliling: %d, Luas: %d\n", keliling, luas)
}
