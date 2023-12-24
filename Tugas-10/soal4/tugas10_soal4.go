package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	var phones = []string{}
	tambahPhone("Xiaomi", &phones)
	tambahPhone("Asus", &phones)
	tambahPhone("IPhone", &phones)
	tambahPhone("Samsung", &phones)
	tambahPhone("Oppo", &phones)
	tambahPhone("Realme", &phones)
	tambahPhone("Vivo", &phones)

	sort.Strings(phones)

	for index, phone := range phones {
		fmt.Printf("%d. %s\n", index+1, phone)
		time.Sleep(time.Second * 1)
	}
}

func tambahPhone(newPhone string, phones *[]string) {
	*phones = append(*phones, newPhone)
}
