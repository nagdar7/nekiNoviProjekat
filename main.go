package main

import (
	"fmt"
	"nekiNoviProjekat/bzvz"
	"strconv"
)

func printOneToN(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func swap(x int, y string) (a int, b string) {
	a, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	b = strconv.Itoa(x)
	return
}

func main() {

	bzvz.Bzvz()

	a, b := 5, "4"

	a, b = swap(a, b)

	fmt.Println(a, b)

	printOneToN(20)

}
