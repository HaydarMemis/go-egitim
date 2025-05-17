package main

import "fmt"

func main() {

	var isim string
	var yas int
	fmt.Print("Adinizi Giriniz: ")
	fmt.Scan(&isim)

	fmt.Print("Yasinizi Giriniz: ")
	fmt.Scan(&yas)

	fmt.Printf("Adinz: %s Yasiniz: %d\n", isim, yas)

}
