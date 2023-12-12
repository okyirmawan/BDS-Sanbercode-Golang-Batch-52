package main

import (
	"fmt"
	"strconv"
)

func main() {

	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)

	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang + lebar)
	var luasSegitiga int = alas * tinggi / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	var indexNilaiJohn string
	var indexNilaiDoe string

	if nilaiJohn >= 80 {
		indexNilaiJohn = "indeksnya A"
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		indexNilaiJohn = "indeksnya B"
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		indexNilaiJohn = "indeksnya C"
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		indexNilaiJohn = "indeksnya D"
	} else {
		indexNilaiJohn = "indeksnya E"
	}

	if nilaiDoe >= 80 {
		indexNilaiDoe = "indeksnya A"
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		indexNilaiDoe = "indeksnya B"
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		indexNilaiDoe = "indeksnya C"
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		indexNilaiDoe = "indeksnya D"
	} else {
		indexNilaiDoe = "indeksnya E"
	}

	fmt.Println("Indeks Nilai John:", indexNilaiJohn)
	fmt.Println("Indeks Nilai Doe:", indexNilaiDoe)

	// soal 3
	var tanggal = 1
	var bulan = 10
	var tahun = 1992
	var bulanID string

	switch {
	case bulan == 1:
		bulanID = "Januari"
	case bulan == 2:
		bulanID = "Februari"
	case bulan == 3:
		bulanID = "Maret"
	case bulan == 4:
		bulanID = "April"
	case bulan == 5:
		bulanID = "Mei"
	case bulan == 6:
		bulanID = "Juni"
	case bulan == 7:
		bulanID = "Juli"
	case bulan == 8:
		bulanID = "Agustus"
	case bulan == 9:
		bulanID = "September"
	case bulan == 10:
		bulanID = "Oktober"
	case bulan == 11:
		bulanID = "November"
	case bulan == 12:
		bulanID = "Desember"
	}

	fmt.Println(tanggal, bulanID, tahun)

	// soal 4
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generasi Y (Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Generasi Z")
	}
}
