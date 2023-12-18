package main

import (
	"fmt"
	"math"
)

func main() {
	// soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64
	jariJari := 10.0

	hitungLuasLingkaran(&luasLingkaran, jariJari)
	hitungKelilingLingkaran(&kelilingLingkaran, jariJari)

	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luasLingkaran)
	fmt.Printf("Keliling lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, kelilingLingkaran)

	// soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{}
	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	for i, namaBuah := range buah {
		fmt.Printf("%d. %s\n", i+1, namaBuah)
	}

	// soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, item["title"])
		fmt.Printf("   duration : %s\n", item["duration"])
		fmt.Printf("   genre : %s\n", item["genre"])
		fmt.Printf("   year : %s\n", item["year"])
		fmt.Println()
	}
}

// soal 1
func hitungLuasLingkaran(luas *float64, jariJari float64) {
	*luas = math.Pi * jariJari * jariJari
}

func hitungKelilingLingkaran(keliling *float64, jariJari float64) {
	*keliling = 2 * math.Pi * jariJari
}

// soal 2
func introduce(sentence *string, name string, gender string, profession string, age string) {
	var salutation string
	if gender == "laki-laki" {
		salutation = "Pak"
	} else if gender == "perempuan" {
		salutation = "Bu"
	}

	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", salutation, name, profession, age)
}

// soal 3
func tambahBuah(buah *[]string, buahBaru string) {
	*buah = append(*buah, buahBaru)
}

// soal 4
func tambahDataFilm(title string, duration string, genre string, year string, listFilm *[]map[string]string) {
	film := map[string]string{
		"genre":    genre,
		"duration": duration,
		"year":     year,
		"title":    title,
	}
	*listFilm = append(*listFilm, film)
}
