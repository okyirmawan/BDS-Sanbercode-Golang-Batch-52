package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

func kelilingSegitigaSamaSisi(sisi int, isSamaSisi bool) (string, error) {
	defer recoverError()

	if isSamaSisi {
		if sisi > 0 {
			return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, sisi*3), nil
		} else {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	} else {
		if sisi > 0 {
			return strconv.Itoa(sisi), nil
		} else {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}
}

func recoverError() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error", message)
	}
}
