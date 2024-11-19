package main

import "fmt"

func main() {
	var a int
	var teks string
	fmt.Scan(&a)
	teks = "Bukan Positif"
	if a > 0 {
		teks = "positif"
	}
	fmt.Println(teks)
}
