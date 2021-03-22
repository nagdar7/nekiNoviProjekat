package main

import "fmt"

func izmjeniMe(b *int) {
	c := 3
	b = &c
	*b = 7
}

func main() {
	c := 4
	izmjeniMe(&c)
	fmt.Println(c)
}
