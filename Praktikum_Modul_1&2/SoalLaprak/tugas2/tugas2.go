package main

import "fmt"

func main() {
	var (
		nama, nim, class string
	)
	fmt.Print("Masukan nama, NIM, dan kelas")
	fmt.Scan(&nama, &nim, &class)
	fmt.Println("Perkenalkan saya adalah ", nama, " salah satu mahasiswa Prodi S1-IF dari kelas ", class, " dengan NIM ", nim)
}
