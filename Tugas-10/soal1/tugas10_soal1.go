package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer endProcess("Golang Backend Development", 2021)
	fmt.Println("Start Program")
}

func endProcess(kalimat string, tahun int) {
	fmt.Println(kalimat, strconv.Itoa(tahun))
}
