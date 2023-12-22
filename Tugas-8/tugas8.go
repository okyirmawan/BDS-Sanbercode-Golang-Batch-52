package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// soal 1
	fmt.Println("============================")
	fmt.Println("Soal 1")
	fmt.Println("============================")
	var bangunDatar hitungBangunDatar

	bangunDatar = segitigaSamaSisi{10, 10}
	fmt.Println("===== Segitiga Sama Sisi =====")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())
	fmt.Println()

	bangunDatar = persegiPanjang{15, 5}
	fmt.Println("===== Persegi Panjang =====")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())
	fmt.Println()

	var bangunRuang hitungBangunRuang

	bangunRuang = tabung{14, 10}
	fmt.Println("===== Tabung =====")
	fmt.Println("volume		:", bangunRuang.volume())
	fmt.Println("luas permukaan	:", bangunRuang.luasPermukaan())
	fmt.Println()

	bangunRuang = balok{4, 5, 6}
	fmt.Println("===== Balok =====")
	fmt.Println("volume		:", bangunRuang.volume())
	fmt.Println("luas permukaan	:", bangunRuang.luasPermukaan())
	fmt.Println()

	// soal 2
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 2")
	fmt.Println("============================")
	var phoneInfo phoneInfo = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze, Mystic White, Mystic Black"},
	}

	fmt.Println(phoneInfo.showInfo())

	// soal 3
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 3")
	fmt.Println("============================")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// soal 4
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 4")
	fmt.Println("============================")

	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)
	gabunganAngka := append(angkaPertama, angkaKedua...)

	total := sum(gabunganAngka)

	result := fmt.Sprintf("%s%s = %d", prefix.(string), formatNumbers(gabunganAngka), total)
	fmt.Println(result)

}

// Segitiga Sama Sisi
type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

// Persegi Panjang
type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
	return 2 * (pp.panjang + pp.lebar)
}

// Tabung
type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return (2 * math.Pi * t.jariJari) * (t.jariJari + t.tinggi)
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * ((b.panjang * b.lebar) + (b.panjang * b.tinggi) + (b.lebar * b.tinggi)))
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p phone) showInfo() string {
	colors := strings.Join(p.colors, ", ")
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.name, p.brand, p.year, colors)
}

type phoneInfo interface {
	showInfo() string
}

func luasPersegi(sisi int, isSamaSisi bool) interface{} {
	if sisi == 0 && isSamaSisi {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if isSamaSisi {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, sisi*sisi)
	} else if isSamaSisi == false && sisi > 0 {
		return sisi
	}

	return nil
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func formatNumbers(nums []int) string {
	formattedNumbers := make([]string, len(nums))

	for i, num := range nums {
		formattedNumbers[i] = fmt.Sprint(num)
	}

	return strings.Join(formattedNumbers, " + ")
}
