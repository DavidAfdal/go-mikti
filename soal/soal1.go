package soal

import "fmt"

func Soal1() {

	fmt.Println("=========== SOAL 1 ===========")

	var skorLumbaLumba = []int{96, 108, 89}
	var skorKoala = []int{88, 91, 110}

	rataRataLumbaLumba := hitungRataRata(skorLumbaLumba)
	rataRataKoala := hitungRataRata(skorKoala)

	fmt.Println("Rata - rata skor lumba - lumba : ", rataRataLumbaLumba)
	fmt.Println( "Rata - rata skor koala : ", rataRataKoala)


	if rataRataLumbaLumba > rataRataKoala {
		fmt.Println("Tim lumba-lumba menang")
	} else if rataRataLumbaLumba < rataRataKoala {
		fmt.Println("Tim koala menang")
	} else {
		fmt.Println("Kedua Tim Seri")
	}

}
func Soal1Bonus1() {

	fmt.Println("=========== Bonus 1 ===========")

	var skorLumbaLumba = []int{97, 112, 101}
	var skorKoala = []int{109, 95, 123}

	rataRataLumbaLumba := hitungRataRata(skorLumbaLumba)
	rataRataKoala := hitungRataRata(skorKoala)

	fmt.Println("Rata - rata skor lumba - lumba : ", rataRataLumbaLumba)
	fmt.Println( "Rata - rata skor koala : ", rataRataKoala)

	if  rataRataLumbaLumba == rataRataKoala {
		fmt.Println("Kedua Tim Seri")
		return
	}

	if rataRataLumbaLumba < 100 || rataRataKoala < 100 { 
		fmt.Println("tidak ada tim yang memenangkan trofi karena tidak memenuhi skor minimal 100")
		return
	} 

	if rataRataLumbaLumba > rataRataKoala {
		fmt.Println("Tim lumba-lumba menang")
	} else {
		fmt.Println("Tim Kaoala menang")
	} 
}
func Soal1Bonus2() {

	fmt.Println("=========== Bonus 2 ===========")

	var skorLumbaLumba = []int{97, 112, 101}
	var skorKoala = []int{109, 95, 106}


	rataRataLumbaLumba := hitungRataRata(skorLumbaLumba)
	rataRataKoala := hitungRataRata(skorKoala)

	fmt.Println("Rata - rata skor lumba - lumba : ", rataRataLumbaLumba)
	fmt.Println( "Rata - rata skor koala : ", rataRataKoala)

	if rataRataLumbaLumba < 100 || rataRataKoala < 100 { 
		fmt.Println("tidak ada tim yang memenangkan trofi karena tidak memenuhi persyaratan")
		return
	} 

	if rataRataLumbaLumba > rataRataKoala {
		fmt.Println("Tim lumba-lumba menang")
	} else if rataRataLumbaLumba < rataRataKoala {
		fmt.Println("Tim Koala menang")
	} else {
		fmt.Println("Kedua Tim Seri")
	}

}

func hitungRataRata(data []int) int {
	var total int = 0
	var rataRata int
	for _, v := range data {
		total += v
	}
	rataRata = total / len(data)
	return rataRata
}