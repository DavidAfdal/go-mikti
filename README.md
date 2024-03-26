# Mikti Challenge 2
Study case menggunakan bahasa pemrograman Golang


## Main.go

Running :

```bash
go run main.go

```

```go
package main

import "github.com/davidafdal/go-mikti/soal"

func main() {
	soal.Soal1()
	soal.Soal1Bonus1()
	soal.Soal1Bonus2()
	soal.Soal2()
}
```



## Soal 1


Ada dua tim senam, Tim Lumba-lumba dan Tim Koala. Mereka bertanding satu sama lain sebanyak 3 kali. Pemenang dengan skor rata-rata tertinggi memenangkan trofi!
Tugas kamu:
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





