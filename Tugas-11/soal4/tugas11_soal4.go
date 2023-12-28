package main

import "fmt"

func luasPersegiPanjang(panjang int, lebar int, ch chan int) {
	ch <- panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int, ch chan int) {
	ch <- 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int, ch chan int) {
	ch <- panjang * lebar * tinggi
}

func main() {
	panjang := 10
	lebar := 7
	tinggi := 5

	chLuas := make(chan int)
	go luasPersegiPanjang(panjang, lebar, chLuas)

	chKeliling := make(chan int)
	go kelilingPersegiPanjang(panjang, lebar, chKeliling)

	chVolume := make(chan int)
	go volumeBalok(panjang, lebar, tinggi, chVolume)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-chLuas:
			fmt.Printf("Luas Persegi Panjang %d\n", luas)
		case keliling := <-chKeliling:
			fmt.Printf("Keliling Persegi Panjang %d\n", keliling)
		case volume := <-chVolume:
			fmt.Printf("Volume Balok %d\n", volume)

		}
	}
}
