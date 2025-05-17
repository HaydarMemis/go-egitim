package main

import "fmt"

func main() {

	//kosulsuz switch
	/*
		sayi := 103

		switch sayi {
		case 5:
			fmt.Println("Sayi 5e Esittir")
		case 10:
			fmt.Println("Sayi 10a Esittir")
		default:
			fmt.Println("Tanimsiz")
		}

	*/

	/*
		harf := "h"
		switch harf {
		case "a", "e", "i":
			fmt.Println("Sesli Harf")
		default:
			fmt.Println("Sessiz Harf")
		}

	*/

	//KOSULLU SWITCH switch yaninda tanimli bi sey yok

	/*
		sayi := 5

		switch {
		case sayi == 5:
			fmt.Println("Sayi 5e Esittir")
			fallthrough //iki case'e de bakiyor cunku 5 hem 10dan kucuk hem de tanimlanan degerdir
		case sayi < 10:
			fmt.Println("Sayi 10dan Buyuktur")
		default:
			fmt.Println("Tanimsiz")

		}

	*/

	switch sayi := 42; {
	case sayi >= 10 && sayi <= 10:
		fmt.Println("Sayi 0dan buyuk 10dan kucuk yada esit")
	case sayi >= 11 && sayi <= 30:
		fmt.Println("Sayi 11den buyuk yada esit yada 30dan kucuk yada esit")
	default:
		fmt.Println("Tanimsiz Aralik")
	}

}
