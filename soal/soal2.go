package soal

import "fmt"

func Soal2() {
	fmt.Println("=========== SOAL 2 ===========")

	var markHigherBMI bool

	massaMark := 78
	tinggiMark := 1.69


	massaJohn := 92
	tinggiJohn := 1.95

	bmiMark := hitungBMI(massaMark, tinggiMark)
	bmiJohn := hitungBMI(massaJohn, tinggiJohn)

	fmt.Printf("BMI Mark : %.2f \n",  bmiMark)
	fmt.Printf("BMI John : %.2f \n",  bmiJohn)


	if bmiMark > bmiJohn {
		markHigherBMI = true
	} else {
		markHigherBMI = false
	}
	
   fmt.Println("Apakah BMI mark lebih tinggi dari john ? ", markHigherBMI)

}

func hitungBMI(massa int, tinggi float64) float64 {
	bmi := float64(massa) / (tinggi * tinggi)

	return bmi
}