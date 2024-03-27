package soal

import "fmt"

func Soal2Data1() {
	fmt.Println("=========== SOAL 2 Data 1===========")

	var markHigherBMI bool

	beratMark := 78
	tinggiMark := 1.69


	beratJohn := 92
	tinggiJohn := 1.95

	bmiMark := hitungBMI(beratMark, tinggiMark)
	bmiJohn := hitungBMI(beratJohn, tinggiJohn)

	fmt.Printf("BMI Mark : %.2f \n",  bmiMark)
	fmt.Printf("BMI John : %.2f \n",  bmiJohn)


	if bmiMark > bmiJohn {
		markHigherBMI = true
	} else {
		markHigherBMI = false
	}
	
   fmt.Println("Apakah BMI mark lebih tinggi dari john ? ", markHigherBMI)

}
func Soal2Data2() {
	fmt.Println("=========== SOAL 2 Data 2 ===========")

	var markHigherBMI bool

	beratMark := 95
	tinggiMark := 1.88


	beratJohn := 85
	tinggiJohn := 1.76

	bmiMark := hitungBMI(beratMark, tinggiMark)
	bmiJohn := hitungBMI(beratJohn, tinggiJohn)

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