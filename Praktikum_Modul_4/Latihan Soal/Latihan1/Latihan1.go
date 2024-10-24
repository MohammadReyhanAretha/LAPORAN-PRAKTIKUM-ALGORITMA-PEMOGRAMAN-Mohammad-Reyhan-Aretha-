package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	fmt.Print("Total Belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("diskon: ")
	fmt.Scan(&diskon)

	diskon = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Print("Total harga belanja setelah diskon: ", diskon)
}
