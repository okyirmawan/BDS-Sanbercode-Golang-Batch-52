package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn) // halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(title string, duration string, genre string, year string) {
		film := map[string]string{
			"genre": genre,
			"jam":   duration,
			"tahun": year,
			"title": title,
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}

// soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// soal 2
func introduce(name string, gender string, profession string, age string) (result string) {
	var salutation string
	if gender == "laki-laki" {
		salutation = "Pak"
	} else if gender == "perempuan" {
		salutation = "Bu"
	}

	result = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", salutation, name, profession, age)
	return
}

// soal 3
func buahFavorit(name string, fruits ...string) (result string) {
	for i, v := range fruits {
		fruits[i] = fmt.Sprintf(`"%s"`, v)
	}

	favoriteFruit := strings.Join(fruits, ", ")
	result = fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, favoriteFruit)
	return
}
