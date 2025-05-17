package main

import "fmt"

func main() {

	//fmt.Println("Merhaba\n\tDunya") //\n \t

	fmt.Println("Haydar Memis'in \"Bilgisayari\" ")
	fmt.Println("Malatya" + "Kayisisi" + "Ile" + "Unludur")
	fmt.Println("Malatya", "Kayisisi", "Ile", "Unludur")

	fmt.Println("Lorem Ipsum" +
		"Merhaba Dunya" +
		"Hello World")

	isim := "Haydar"
	soyad := "Memis"
	sehir := "Ankara"

	fmt.Println(isim, soyad, sehir)

	//fmt.Println(len(isim))
	var stringUzunlugu = len(isim)
	fmt.Println(stringUzunlugu)

}
