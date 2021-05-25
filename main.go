package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func print1ToN(n int) {
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

func radSaStringovima() {
	a := "idu  tri konja ulicom"
	b := ""
	for _, s := range a {
		if s == ' ' {
			b += string('_')
		} else {
			b += string(s)
		}
	}
	fmt.Println(b)

	a = "idu  tri konja ulicom"
	bb := strings.Builder{}
	for _, s := range a {
		if s == ' ' {
			bb.WriteString("_")
		} else {
			bb.WriteRune(s)
		}
	}
	fmt.Println(bb.String())

	bbb := strings.Fields(a)
	fmt.Println(strings.Join(bbb, "_"))

	fmt.Println(strings.HasPrefix(a, "idu"))
	fmt.Println(strings.HasSuffix(a, "ulicom"))

	fmt.Println(strings.Index(a, " "))

	fmt.Println(strings.LastIndex(a, " "))

	b = strings.Map(func(z rune) rune {
		if z == ' ' {
			return -1
		}
		return z
	}, a)
	fmt.Println(b)

	r := strings.NewReplacer(" ", "_", "nece", "XD", "idu", "nece")
	b = r.Replace(a)
	ab := r.Replace(b)
	fmt.Println(b, ab)

	b = strings.ReplaceAll(a, " ", "_")
	fmt.Println(b)

	d := strings.Repeat("bla", 7)
	fmt.Println(d)

	f := strings.Split(a, " ")
	fmt.Println(strings.Join(f, "_"))

	e := strings.SplitAfter(a, " ")
	fmt.Println(strings.Join(e, "_"))

	fmt.Println(strings.Title(a))

	fmt.Println(strings.Trim(a, "iodum"))

	fmt.Println(strings.TrimPrefix(a, "i"))

	a = " \n\r" + a + " "
	fmt.Println(a)
	fmt.Println(strings.TrimSpace(a))
}

func CezarSifra(tekst string, kljuc int) string {
	kljuc = kljuc % 26
	if kljuc < 0 {
		kljuc = kljuc + 26
	}
	a := ""
	for _, s := range tekst {
		if s < 'A' || s > 'z' || s < 'a' && s > 'Z' {
			a += string(s)
			continue
		}
		r := s + rune(kljuc)
		if r > 'Z' && s >= 'A' && s <= 'Z' {
			r = r - 26
		}
		if r > 'z' {
			r = r - 26
		}
		a = a + string(r)
	}
	return a
}

func isPalindrom(tekst string) bool {
	tekst = strings.ToLower(tekst)
	a := strings.Fields(tekst)
	b := strings.Join(a, "")
	fmt.Println(b)
	for i := 0; i <= len(b)/2; i++ {
		if b[i] == b[len(b)-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func filterSlova(tekst string) string {
	b := ""
	for _, s := range tekst {
		if s == ' ' || s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
			b = b + string(s)
		} else {
			continue
		}
	}
	return b
}

func filterSlovaMap(tekst string) string {
	b := strings.Map(func(s rune) rune {
		if s == ' ' || s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
			return s
		}
		return -1
	}, tekst)
	return b
}

// //func sortiraj(lista []string) []string {
// 	for i, s := range lista {

// 	}
// 	return nil
// }

func porediString(a, b string) int {
	a, b = strings.ToUpper(a), strings.ToUpper(b)
	maxlen := len(a)
	if len(a) > len(b) {
		maxlen = len(b)
	}
	for i := 0; i < maxlen; i++ {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(b) < len(a) {
		return 1
	}
	return 0
}

func sortiraj(tekst []string) []string {
	for i := 0; i < len(tekst)-1; i++ {
		for j := i; j < len(tekst); j++ {
			a := porediString(tekst[i], tekst[j])
			if a > 0 {
				tekst[i], tekst[j] = tekst[j], tekst[i]
			}
		}
	}
	return tekst
}

func sortInt(XD []int) []int {
	for i := 0; i < len(XD)-1; i++ {
		for j := i + 1; j < len(XD); j++ {
			if XD[i] > XD[j] {
				XD[i], XD[j] = XD[j], XD[i]
			}
		}
	}
	return XD
}

func najduzaRijec(recenica string) string {
	recenica = filterSlovaMap(recenica)
	c := strings.Fields(recenica)
	if len(c) == 0 {
		return ""
	}
	max := 0
	for i, _ := range c {
		if len(c[max]) < len(c[i]) {
			max = i
		}
	}
	return c[max]
}

func isLongestPalindrom(tekst string) string {
	tekst = strings.ToLower(tekst)
	a := strings.Fields(tekst)
	b := strings.Join(a, "")
	max := 0
	maxstr := ""
	maxi := 0
	maxj := 0
	for i := 0; i <= len(b)/2; i++ {
		for j := len(b) - 1; j > i; j-- {
			// ako su i i j izmedju maxi i maxj, continue
			if i >= maxi && i <= maxj && j >= maxi && j <= maxj {
				continue
			}
			if b[i] == b[j] && isPalindrom(b[i:j+1]) {
				if j+1-i > max {
					max = j + 1 - i
					maxstr = b[i : j+1]
					maxi = i
					maxj = j
				}
			}
		}
	}
	return maxstr
}

func konvert(broj int) []string {
	c := []string{}
	if broj < 0 {
		return c
	}
	for broj > 0 {
		a := broj % 10
		b := strconv.Itoa(a)
		c = append(c, b)
		broj = broj / 10
	}
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-1-i] = c[len(c)-1-i], c[i]
	}
	return c
}

func najveci(lista []int) (int, int) {
	max := 0
	maxIdx := -1
	for i, s := range lista {
		if s > max || maxIdx < 0 {
			max = s
			maxIdx = i
		}
	}
	return max, maxIdx
}

func spajanjeListi(lista1 []interface{}, lista2 []interface{}) []interface{} {
	c := []interface{}{}
	minLen := len(lista1)
	if len(lista2) < len(lista1) {
		minLen = len(lista2)
	}
	for i := 0; i < minLen; i++ {
		c = append(c, lista1[i])
		c = append(c, lista2[i])
	}
	if minLen == len(lista1) && len(lista2) != minLen {
		c = append(c, lista2[minLen:]...)
	}
	if minLen == len(lista2) && len(lista1) != minLen {
		c = append(c, lista1[minLen:]...)
	}
	return c
}

func brisanjeElementa(lista []int, idx int) []int {
	c := []int{}
	for i := 0; i < len(lista); i++ {
		if i != idx {
			c = append(c, lista[i])
		}
	}
	return c
}

func brisanjeElementa2(lista []int, idx int) []int {
	if idx < 0 || idx >= len(lista) {
		return lista
	}
	return append(lista[:idx], lista[idx:]...)
}

func filterNeparnih(niz []int) []int {
	niz2 := []int{}
	for _, s := range niz {
		if s%2 == 0 {
			niz2 = append(niz2, s)
		}
	}
	return niz2
}

func filter(niz []int, check func(int, int) bool) []int {
	niz2 := []int{}
	for i, s := range niz {
		if check(i, s) {
			niz2 = append(niz2, s)
		}
	}
	return niz2
}

var stack = []interface{}{}

var stackMaxElem = 5

func push(elem interface{}) error {
	if len(stack) >= stackMaxElem {
		return errors.New("stack overflow")
	}
	stack = append(stack, elem)
	return nil
}

func pop() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("stack empty")
	}
	a := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return a, nil
}

func sortNiz(XD []int) []int {
	for i := 0; i < len(XD)-1; i++ {
		for j := i + 1; j < len(XD); j++ {
			if XD[i] > XD[j] {
				XD[i], XD[j] = XD[j], XD[i]
			}
		}
	}
	return XD
}

func main() {
	//radSaStringovima()
	//a := "Ana voli milovana"
	//b := CezarSifra(a, 13)
	//c := CezarSifra(b, 13)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(isPalindrom("ana voli milovana"))
	//fmt.Println(isPalindrom(" ana voli milovana"))
	//fmt.Println(isPalindrom("ana ne voli milovana"))
	// fmt.Println(filterSlovaMap("Banana//// znao sam to"))
	// c := "Banana znao s"
	// fmt.Println(sortiraj(strings.Fields(c)))
	// g := []int{3, 6, 89, 4, 2, 56, 8, 90, 4, 2, 5, 8, 9}
	// // fmt.Println(sortInt(g))
	// // orderBy := "ltog"
	// orderBy := "gtol"

	// d := strings.Fields(c)
	// fmt.Println(d)

	// sort.Slice(g, func(i, j int) bool {
	// 	if orderBy == "ltog" {
	// 		return g[i] < g[j]
	// 	} else if orderBy == "gtol" {
	// 		return g[i] > g[j]
	// 	}
	// 	return true
	// })
	// fmt.Println(g)

	// c := "Ana voli Milovana 2 5 Ana voli Milovana Ana voli Milovana Ana voli Milovana"
	// b := " "
	// d := "ana b..."
	// fmt.Println(najduzaRijec(c))
	// fmt.Println(najduzaRijec(b))
	// fmt.Println(najduzaRijec(d))

	// fmt.Println(isLongestPalindrom(c))

	// fmt.Println(c[0:3])

	// fmt.Println(konvert(12345))

	// fmt.Println(najveci([]int{2, 4, 7, 20, 3, 7, 8, 5}))
	// fmt.Println(najveci([]int{}))
	// fmt.Println(najveci(nil))

	// fmt.Println(spajanjeListi([]interface{}{2, 4, 7, 20, 3, 7, 8, 5}, []interface{}{"dva", "tri", "pet", "osam"}))
	// fmt.Println(spajanjeListi([]interface{}{}, []interface{}{}))
	// fmt.Println(spajanjeListi(nil, nil))
	// fmt.Println(spajanjeListi([]interface{}{2, 4, 7, 20, 3, 7, 8, 5}, []interface{}{}))

	// fmt.Println(brisanjeElementa2([]int{2, 5, 6, 7, 8}, -1))
	// fmt.Println(brisanjeElementa2([]int{2, 5, 6, 7, 8}, 6))

	fmt.Println(filterNeparnih([]int{2, 4, 7, 20, 3, 7, 8, 5}))
	fmt.Println(filter([]int{2, 4, 7, 20, 3, 7, 8, 5}, func(i, val int) bool {
		return val > 3
	}))

	fmt.Println(pop())
	push(3)
	push("5")
	push(7)
	push("bla")
	push(2)
	fmt.Println(push(5))
	fmt.Println(stack)
	fmt.Println(pop())
	fmt.Println(stack)

	fmt.Println(sortNiz([]int{2, 4, 7, 20, 3, 7, 8, 5}))
}
