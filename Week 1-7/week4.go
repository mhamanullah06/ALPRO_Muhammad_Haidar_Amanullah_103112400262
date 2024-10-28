package main

import "fmt"

func main() {
	var panjang, lebar float64

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Printf("\nInformasi Persegi Panjang:\n")
	fmt.Printf("Panjang: %.2f\n", panjang)
	fmt.Printf("Lebar: %.2f\n", lebar)
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}
