package main

import "fmt"

type dikdortgen struct {
	kisakenar, uzunkenar float64
}

func (a dikdortgen) alanhesapla() float64 {
	return a.kisakenar * a.uzunkenar
}

func (b dikdortgen) cevrehesapla() float64 {
	return 2 * (b.kisakenar + b.uzunkenar)
}

func main() {
	r := dikdortgen{kisakenar: 5, uzunkenar: 7}
	fmt.Println(r.alanhesapla())
	fmt.Println(r.cevrehesapla())
}
