# Mikti Challenge 2
Study case menggunakan bahasa pemrograman Golang


## Main.go

Code :

```go
package main

import "github.com/davidafdal/go-mikti/soal"

func main() {
	soal.Soal1()
	soal.Soal1Bonus1()
	soal.Soal1Bonus2()
	soal.Soal2Data1()
	soal.Soal2Data2()
}
```

Running :

```bash
go run main.go

```


## Soal 1


Ada dua tim senam, Tim Lumba-lumba dan Tim Koala. Mereka bertanding satu sama lain sebanyak 3 kali. Pemenang dengan skor rata-rata tertinggi memenangkan trofi!
Tugas :
- Hitung skor rata-rata untuk setiap tim, menggunakan data uji di bawah ini.
```bash
Data Uji:
Data 1: Lumba-lumba mendapat skor 96, 108 dan 89. Koala mendapat skor 88, 91 dan 110.
Data Bonus 1: Lumba-lumba mendapat skor 97, 112 dan 101. Koala mendapat skor 109, 95 dan 123.
Data Bonus 2: Lumba-lumba mendapat skor 97, 112 dan 101. Koala mendapat skor 109, 95 dan 106
```
- Bandingkan skor rata-rata tim untuk menentukan pemenang kompetisi, dan cetak ke konsol. Jangan lupa bahwa bisa ada hasil seri, jadi uji juga untuk itu (seri berarti mereka memiliki skor rata-rata yang sama).
- Bonus 1: Sertakan persyaratan untuk skor minimum 100. Dengan aturan ini, sebuah tim hanya menang jika memiliki skor lebih tinggi dari tim lain, dan pada saat yang sama skor minimal 100 poin. Petunjuk: Gunakan operator logika untuk menguji skor minimum, serta beberapa blok else-if.
- Bonus 2: Skor minimum juga berlaku untuk seri! Jadi seri hanya terjadi ketika kedua tim memiliki skor yang sama dan keduanya memiliki skor lebih besar atau sama dengan 100 poin. Jika tidak, tidak ada tim yang memenangkan trofi.

Code : 
```go
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
```

Output :
```bash

=========== SOAL 1 ===========
Rata - rata skor lumba - lumba :  97
Rata - rata skor koala :  96
Tim lumba-lumba menang
=========== Bonus 1 ===========
Rata - rata skor lumba - lumba :  103
Rata - rata skor koala :  109
Tim Kaoala menang
=========== Bonus 2 ===========
Rata - rata skor lumba - lumba :  103
Rata - rata skor koala :  103
Kedua Tim Seri

```

## Soal 2

Mark dan John sedang mencoba membandingkan BMI (Body Mass Index) mereka, yang dihitung menggunakan rumus: BMI = massa / tinggi ** 2 = massa / (tinggi * tinggi) (massa dalam kg dan tinggi dalam meter).

Tugas :
- Simpan massa dan tinggi badan Mark dan John dalam variabel.
- Hitung kedua BMI mereka menggunakan rumus (kamu bahkan dapat menerapkan kedua versi)
- Buat variabel Boolean 'markHigherBMI' yang berisi informasi tentang apakah Mark memiliki BMI lebih tinggi dari John.

Data Uji:
```bash
Data 1: Berat Mark 78 kg dan tinggi 1.69 m. Berat John 92 kg dan tinggi 1.95 m.
Data 2: Berat Mark 95 kg dan tinggi 1.88 m. Berat John 85 kg dan tinggi 1.76 m.
```

Code :

```go
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
```

Output :

```bash 
=========== SOAL 2 Data 1===========
BMI Mark : 27.31
BMI John : 24.19
Apakah BMI mark lebih tinggi dari john ?  true
=========== SOAL 2 Data 2 ===========
BMI Mark : 26.88
BMI John : 27.44
Apakah BMI mark lebih tinggi dari john ?  false
```









