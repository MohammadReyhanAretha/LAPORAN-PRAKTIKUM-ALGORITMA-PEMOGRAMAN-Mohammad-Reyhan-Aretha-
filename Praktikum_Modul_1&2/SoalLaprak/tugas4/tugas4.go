package main

import "fmt"

func main() {
	var c, f float64
	fmt.Scan(&f)
	c = (f - 32) * 5 / 9
	fmt.Print(c)
}