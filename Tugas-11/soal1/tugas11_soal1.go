package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func printPhone(index int, phone string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("%d. %s\n", index, phone)
}

func main() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)
	var wg sync.WaitGroup

	for i, phone := range phones {
		wg.Add(1)
		go printPhone(i+1, phone, &wg)
		wg.Wait()
	}
}
