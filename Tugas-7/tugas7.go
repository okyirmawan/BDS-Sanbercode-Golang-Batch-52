package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	fmt.Println("============================")
	fmt.Println("Soal 1")
	fmt.Println("============================")

	type Buah struct {
		Nama, Warna string
		AdaBijinya  bool
		harga       int
	}

	buah1 := Buah{"Nanas", "Kuning", false, 9000}
	buah2 := Buah{"Jeruk", "Oranye", true, 8000}
	buah3 := Buah{"Semangka", "Hijau & Merah", true, 10000}
	buah4 := Buah{"Pisang", "Kuning", false, 5000}

	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah3)
	fmt.Println(buah4)

	// soal 2
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 2")
	fmt.Println("============================")

	segitiga := Segitiga{Alas: 5, Tinggi: 6}
	fmt.Println(segitiga.Luas())

	persegi := Persegi{Sisi: 10}
	fmt.Println(persegi.Luas())

	persegiPanjang := PersegiPanjang{Panjang: 10, Lebar: 5}
	fmt.Println(persegiPanjang.Luas())

	// soal 3
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 3")
	fmt.Println("============================")
	iphone := Phone{
		Name:   "iPhone",
		Brand:  "Apple",
		Year:   2023,
		Colors: []string{},
	}

	iphone.addColor("Silver")
	iphone.addColor("Space Gray")
	iphone.addColor("Gold")

	fmt.Printf("Updated Colors %s(%s) %d: %s\n", iphone.Name, iphone.Brand, iphone.Year, strings.Join(iphone.Colors, ", "))

	// soal 3
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 4")
	fmt.Println("============================")

	var dataFilm = []Movie{}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, item.Title)
		fmt.Printf("   duration : %d\n", item.Duration)
		fmt.Printf("   genre : %s\n", item.Genre)
		fmt.Printf("   year : %d\n", item.Year)
		fmt.Println()
	}
}

// soal 2
type Segitiga struct {
	Alas, Tinggi int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (s Segitiga) Luas() float64 {
	return 0.5 * float64(s.Alas) * float64(s.Tinggi)
}

func (p Persegi) Luas() int {
	return p.Sisi * p.Sisi
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

// soal 3
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (phone *Phone) addColor(newColor string) {
	phone.Colors = append(phone.Colors, newColor)
}

// soal 4
type Movie struct {
	Title, Genre   string
	Duration, Year int
}

func tambahDataFilm(title string, duration int, genre string, year int, listFilm *[]Movie) {
	film := Movie{
		Title:    title,
		Genre:    genre,
		Duration: duration,
		Year:     year,
	}

	*listFilm = append(*listFilm, film)
}
