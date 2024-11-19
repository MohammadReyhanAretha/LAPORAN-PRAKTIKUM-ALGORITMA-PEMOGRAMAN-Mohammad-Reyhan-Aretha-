package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Scan(&a, &b)
	if b != 0 {
		hasil = a / b
		fmt.Println("hasilnya adalah", hasil)
	}
}
