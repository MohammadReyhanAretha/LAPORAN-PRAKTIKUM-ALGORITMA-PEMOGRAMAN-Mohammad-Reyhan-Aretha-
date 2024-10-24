import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukan berat badan (kg): ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}