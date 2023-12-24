package main

import (
	"fmt"
	"math"
)

func main() {
	lingkaran1 := lingkaran{7}
	fmt.Println("Luas Lingkaran 1 :", lingkaran1.luasLingkaran())
	fmt.Println("Keliling Lingkaran 1 :", lingkaran1.kelilingLingkaran())
	fmt.Println()

	lingkaran2 := lingkaran{10}
	fmt.Println("Luas Lingkaran 2 :", lingkaran2.luasLingkaran())
	fmt.Println("Keliling Lingkaran 2 :", lingkaran2.kelilingLingkaran())
	fmt.Println()

	lingkaran3 := lingkaran{15}
	fmt.Println("Luas Lingkaran 3 :", lingkaran3.luasLingkaran())
	fmt.Println("Keliling Lingkaran 3 :", lingkaran3.kelilingLingkaran())
	fmt.Println()
}

type lingkaran struct {
	jarijari float64
}

func (l lingkaran) luasLingkaran() float64 {
	luas := math.Pi * math.Pow(l.jarijari, 2)
	return math.Round(luas)
}

func (l lingkaran) kelilingLingkaran() float64 {
	keliling := 2 * math.Pi * l.jarijari
	return math.Round(keliling)
}
