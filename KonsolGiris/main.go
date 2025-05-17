package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*
		tarayici := bufio.NewScanner(os.Stdin) //scanner objesi

		fmt.Print("Bir seyler yaziniz:")

		tarayici.Scan()

		veriGirisi := tarayici.Text()

		fmt.Printf("Bunu Yazdiniz: %s\n", veriGirisi)

	*/

	//int icin

	tarayici := bufio.NewScanner(os.Stdin)

	fmt.Print("Hangi Yil dogdunuz: ")

	tarayici.Scan()

	verigirisi, _ := strconv.ParseInt(tarayici.Text(), 10, 64)

	fmt.Printf("Su anda %d yasindasin\n", 2025-verigirisi)

}
