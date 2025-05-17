package main

import "fmt"

func main() {

	/*var sayi1 = 25

	fmt.Printf("Benim Yas: %d\n", sayi1) //int de %d &b olursa binary yazar
	*/

	/*var sayi2 float32 = 3.14159
	fmt.Printf("PI Sayisi: %.2f\n", sayi2) //.d virgulden sonra 2 sayi */

	/*var dogruMu bool = true
	fmt.Printf("%t\n", dogruMu)*/

	// karakterler %cpointerlar %p str ise &s

	/*var isim string = "haydar"
	var sehir = "ankara"
	var yas = 26
	fmt.Printf("benim adim %s %d yasindayim ve %s da yasiyorum.\n", isim, yas, sehir)
	*/

	//%T veri tipi belirlemek/ogrenmek icin

	/*
		var sayi1 int = 54
	*/

	/*
		var sayi2 float32 = 5.543
		var yanlisMi bool = false
		var str string = "Ebrar"

		fmt.Printf("%T %T %T %T\n", sayi1, sayi2, yanlisMi, str)
	*/

	var mesaj string = "Merhaba Dunya"
	var yil int = 2025

	var tumMesaj = fmt.Sprintf("%d yilindan %s", yil, mesaj) // burda degisken olarak tanimlamis oldun
	fmt.Println(tumMesaj)

}
