package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scaner := bufio.NewScanner(os.Stdin)

	fmt.Print("Kilonuzu Giriniz: ")
	scaner.Scan()

	kilo, _ := strconv.ParseFloat(scaner.Text(), 64)

	fmt.Print("Boyunuzu Metre olarak Giriniz: ")
	scaner.Scan()

	boy, _ := strconv.ParseFloat(scaner.Text(), 64)

	vki := kilo / (boy * boy)

	fmt.Printf("Vucut kitle endeksiniz: %f\n", vki)

}
