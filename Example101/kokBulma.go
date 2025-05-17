package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("A sayisini giriniz: ")

	scanner.Scan()

	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("B Sayisini giriniz: ")

	scanner.Scan()

	b, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("C Sayisini giriniz: ")

	scanner.Scan()

	c, _ := strconv.ParseFloat(scanner.Text(), 64)

	delta := (math.Pow(b, 2)) - 4*a*c

	kokbir := (-b - (math.Pow(delta, 0.5))) / 2 * a
	kokiki := (-b + (math.Pow(delta, 0.5))) / 2 * a

	fmt.Printf("1.kok: %f\n2.kok: %f\n", kokbir, kokiki)

}
