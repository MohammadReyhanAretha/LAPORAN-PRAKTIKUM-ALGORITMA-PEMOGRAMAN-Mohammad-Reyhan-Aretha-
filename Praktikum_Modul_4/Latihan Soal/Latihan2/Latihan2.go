package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64
	fmt.Print("Nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)
	beratBadan = bmi * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat Badan Anda: %.0f kg\n", beratBadan)
}
