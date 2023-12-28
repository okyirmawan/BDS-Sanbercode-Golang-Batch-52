package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title, Desc, Author string
	ReleaseYear         int
}

func main() {
	books := []Book{
		{"Clean Code", "A Handbook of Agile Software Craftsmanship", "Robert C. Martin", 2008},
		{"Refactoring", "Improving The Design of Existing Code", "Kent Beck", 1999},
		{"Head First Design Pattern", "A Brain-Friendly Guide", "Eric Freeman & Elisabeth Robson", 2004},
		{"Domain Driven Design", "Tackling Complexity in the Heart of Software", "Eric Evans", 2003},
		{"The DevOps Handbook", "How to Create World-Class Agility, Reliability, and Security in Technology Organizations", "Gene Kim", 2016},
	}

	jsonData, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
