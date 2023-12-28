package main

import (
	"encoding/json"
	"fmt"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-12/data"
)

type Phone struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

func main() {
	var jsonString = data.GetJsonPhoneString()
	var phones []Phone

	err := json.Unmarshal([]byte(jsonString), &phones)
	if err != nil {
		fmt.Println("ERR", err)
	}

	fmt.Println("=== Daftar Model Handphone Sony ====")
	number := 1
	for _, phone := range phones {
		if phone.Vendor == "Sony" {
			fmt.Printf("%d. %s - %s\n", number, phone.Vendor, phone.Model)
			number++
		}
	}

}
