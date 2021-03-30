package main

import (
	"fmt"
	"nekiNoviProjekat/bzvz"
	"strconv"
)

var a bool

func swap(x int, y string) (a int, b string) {
	a, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	b = strconv.Itoa(x)
	return
}

func main() {
	a = bzvz.A == 0

	bzvz.Bzvz()

	a, b := 5, "4"

	a, b = swap(a, b)

	fmt.Println(a, b)

}
