package main

import (
	"flag"
	"fmt"
)

func main() {
	panjang := flag.Int("panjang", 10, "Input panjang persegi panjang")
	lebar := flag.Int("lebar", 5, "Input lebar persegi panjang")

	flag.Parse()

	fmt.Println("Luas Persegi Panjang :", luasPersegiPanjang(*panjang, *lebar))
	fmt.Println("Keliling Persegi Panjang :", kelilingPersegiPanjang(*panjang, *lebar))

	// run command
	// go run .\Tugas-10\soal6\tugas10_soal6.go -panjang=50 -lebar=10
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}
