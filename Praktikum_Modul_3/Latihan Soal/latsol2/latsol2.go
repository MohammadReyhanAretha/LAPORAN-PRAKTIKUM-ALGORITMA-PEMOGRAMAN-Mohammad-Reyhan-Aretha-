package main

import "fmt"

func main() {
	var r, v, l float64
	fmt.Print("Jejari: ")
	fmt.Scan(&r)
	v = (4.0 / 3.0) * 3.1415926535 * (r * r * r)
	l = 4 * 3.1415926535 * (r * r)
	fmt.Println("Bola dengan jejari: ", r, "Memiliki Volume: ", v, "dan Luas kulit", l)
}
