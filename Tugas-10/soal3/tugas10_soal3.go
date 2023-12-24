package main

import "fmt"

func main() {
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

func cetakAngka(angka *int) {
	fmt.Println("Total Angka adalah :", *angka)
}

func tambahAngka(tambahAngka int, total *int) {
	*total += tambahAngka
}
