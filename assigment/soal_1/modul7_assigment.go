package main

import "fmt"

func main() {

	var x, y, total int
	fmt.Print("Masukan 2 bilangan yang ingin dijumlahkan :")
	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		total = total + i
	}

	fmt.Println("hasilnya adalah: ", total)
}
