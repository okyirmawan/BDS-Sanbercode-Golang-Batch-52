package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var firstWord = "Bootcamp"
	var secondWord = "Digital"
	var thirdWord = "Skill"
	var fourthWord = "Sanbercode"
	var fifthWord = "Golang"

	concatenatedString := firstWord + " " + secondWord + " " + thirdWord + " " + fourthWord + " " + fifthWord

	fmt.Println(concatenatedString)

	// soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", -1)

	fmt.Println(halo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.ToUpper(string(kataKedua[0])) + kataKedua[1:]

	lastIndexKataKetiga := len(kataKetiga) - 1
	kataKetiga = kataKetiga[:lastIndexKataKetiga] + strings.ToUpper(string(kataKetiga[lastIndexKataKetiga]))
	kataKeempat = strings.ToUpper(kataKeempat)

	gabunganKata := fmt.Sprintf("%s %s %s %s", kataPertama, kataKedua, kataKetiga, kataKeempat)

	fmt.Println(gabunganKata)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	total := angka1 + angka2 + angka3 + angka4

	fmt.Println(total)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimatAngka := `"` + strings.Replace(kalimat, "halo", "Hi", -1) + `"` + " - " + strconv.Itoa(angka)

	fmt.Println(kalimatAngka)
}
