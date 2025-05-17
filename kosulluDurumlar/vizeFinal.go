package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Vize Notunu Giriniz: ")
	scanner.Scan()

	vizenot, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Final Notunu Giriniz: ")
	scanner.Scan()

	finalnot, _ := strconv.ParseFloat(scanner.Text(), 64)

	ortalama := (vizenot * 0.4) + (finalnot * 0.6)

	fmt.Println("ortalama: ", ortalama)

	if ortalama >= 85 && ortalama <= 100 {
		fmt.Println("AA")
	} else if ortalama >= 70 && ortalama < 85 {
		fmt.Println("BB")
	} else if ortalama >= 60 && ortalama < 70 {
		fmt.Println("CC")
	} else if ortalama >= 50 && ortalama < 60 {
		fmt.Println("DD")
	} else if ortalama < 49 && ortalama >= 0 {
		fmt.Println("FF")
	} else {
		fmt.Println("O ile 100 Arasinda bir Deger Olmalidir!")
	}

}
