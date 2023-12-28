package main

import (
	"encoding/json"
	"fmt"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-12/data"
)

func main() {
	var jsonString = data.GetJsonPhoneString()
	var phones []map[string]string

	err := json.Unmarshal([]byte(jsonString), &phones)
	if err != nil {
		fmt.Println("ERR", err)
	}

	fmt.Println("=== Daftar Model Handphone Samsung ====")
	number := 1
	for _, phone := range phones {
		if phone["vendor"] == "Samsung" {
			fmt.Printf("%d. %s - %s\n", number, phone["vendor"], phone["model"])
			number++
		}
	}

}
