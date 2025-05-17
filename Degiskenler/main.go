package main

import "fmt"

func main() {

	var mesaj string = "Merhaba dunya"
	fmt.Println(mesaj)

	var mesaj2 string
	mesaj2 = "Hello World"
	fmt.Println(mesaj2)

	var sayi1 int = 5
	fmt.Println(sayi1)

	var sayi2 float32 = 3.342
	fmt.Println(sayi2)

	var dogruMu bool = true
	fmt.Println(dogruMu)

	var x, y, z int = 3, 4, 5
	fmt.Println(x, y, z)

	var isim = "Haydar" // veri tipi belirtilmek zorunde degil
	fmt.Println(isim)

	var a, b, c = 5.17, false, "Ankara" //farkli veri tiplerini de ayni anda kullanabilirsin
	fmt.Println(a, b, c)

	sehir := "Ankara" // BU Sekilde var kullanmadan da yapabilirsin
	fmt.Println(sehir)

	var sayi10 int = 66
	fmt.Println(sayi10)

	sayi10 = 77
	fmt.Println(sayi10) // yukardaki degisken icerigini degistirmek icin

}
