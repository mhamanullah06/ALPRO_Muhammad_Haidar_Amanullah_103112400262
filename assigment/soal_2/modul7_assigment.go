package main

import "fmt"

func main() {
	var N int

	fmt.Print("Masukan Bilangan Bulat Positif N:")
	fmt.Scanln(&N)

	for i := 1; i <= N; i++ {
		fmt.Println(i * i)
	}

}
