package main

import (
	"fmt"
	"strconv"
)

func main() {

	//TYPE CASTING
	/*
		var toplam int = 984
		var sayi int = 17

		//var sonuc = toplam / sayi --> burasi yerine asagidaki
		var sonuc = float32(toplam) / float32(sayi)

		var s1 = int(sonuc)

		fmt.Println(sonuc)
		fmt.Println(s1)*/

	//TYPE CONVERSION UYUMSUZ VERI TIPI DONUSTURME PAKET CAGIRMAN LAZIM

	/*
		var str = "1"

		//stringden int donusturme

		var sayi, _ = strconv.Atoi(str) //bunu yapmadan string ile int toplanmaz

		fmt.Println(sayi)

		var sonuc = sayi + 7
		fmt.Println(sonuc)*/

	//int ten stringe donusturme

	var sayi = 1

	var str = strconv.Itoa(sayi)

	fmt.Println(str)

	fmt.Println("Nasilsin" + str)

}
