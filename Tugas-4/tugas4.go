package main

import "fmt"

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "-", "Berkualitas")
		} else {
			if i%3 == 0 {
				fmt.Println(i, "-", "I Love Coding")
			} else {
				fmt.Println(i, "-", "Santai")
			}
		}
	}

	// soal 2
	for a := 1; a <= 7; a++ {
		for b := 1; b <= 7; b++ {
			if b <= a {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newKalimat = kalimat[2:]

	fmt.Println(newKalimat)

	// soal 4
	var sayuran = []string{}
	listSayuran := []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	sayuran = append(sayuran, listSayuran...)

	for index, sayur := range sayuran {
		fmt.Printf("%d. %s\n", index+1, sayur)
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)

	}

	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Printf("Volume Balok = %d", volume)

}
