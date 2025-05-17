package main

import "fmt"

/*
	if(kosul){
	islemler
	}
	else{
	islemler
	}
*/

func main() {

	/*
		isim := "mehmet"
		if isim == "haydar" {
			fmt.Println("hosgeldiniz")
		} else if isim == "mehmet" {
			fmt.Println("Hosgeldin Mehmet")
		} else {
			fmt.Println("Yanlis kisi")
		}
	*/

	yas := 19
	fmt.Println("18 yasindan kucukler giremez")

	if yas >= 18 {
		fmt.Println("Hosgeldiniz")
	} else if yas < 18 && yas > 14 {
		fmt.Println("Kosullu giris")
	} else {
		fmt.Println("Giris reddedildi")
	}

}
