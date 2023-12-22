package main

import (
	"fmt"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-9/models"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-9/services"
)

func main() {
	// soal 1
	fmt.Println("============================")
	fmt.Println("Soal 1")
	fmt.Println("============================")
	var bangunDatar models.HitungBangunDatar

	bangunDatar = models.SegitigaSamaSisi{10, 10}
	fmt.Println("===== Segitiga Sama Sisi =====")
	fmt.Println("luas		:", bangunDatar.Luas())
	fmt.Println("keliling	:", bangunDatar.Keliling())
	fmt.Println()

	bangunDatar = models.PersegiPanjang{15, 5}
	fmt.Println("===== Persegi Panjang =====")
	fmt.Println("luas		:", bangunDatar.Luas())
	fmt.Println("keliling	:", bangunDatar.Keliling())
	fmt.Println()

	var bangunRuang models.HitungBangunRuang

	bangunRuang = models.Tabung{14, 10}
	fmt.Println("===== Tabung =====")
	fmt.Println("volume		:", bangunRuang.Volume())
	fmt.Println("luas permukaan	:", bangunRuang.LuasPermukaan())
	fmt.Println()

	bangunRuang = models.Balok{4, 5, 6}
	fmt.Println("===== Balok =====")
	fmt.Println("volume		:", bangunRuang.Volume())
	fmt.Println("luas permukaan	:", bangunRuang.LuasPermukaan())
	fmt.Println()

	// soal 2
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 2")
	fmt.Println("============================")
	var phoneInfo models.PhoneInfo = models.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze, Mystic White, Mystic Black"},
	}

	fmt.Println(phoneInfo.ShowInfo())

	// soal 3
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("Soal 3")
	fmt.Println("============================")
	fmt.Println(services.LuasPersegi(4, true))
	fmt.Println(services.LuasPersegi(8, false))
	fmt.Println(services.LuasPersegi(0, true))
	fmt.Println(services.LuasPersegi(0, false))

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

	total := services.Sum(gabunganAngka)

	result := fmt.Sprintf("%s%s = %d", prefix.(string), services.FormatNumbers(gabunganAngka), total)
	fmt.Println(result)
}
