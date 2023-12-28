package main

import "fmt"

func main() {
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("List Movies:")
	for value := range moviesChannel {
		fmt.Println(value)
	}
}

func getMovies(moviesChannel chan<- string, movies ...string) {
	defer close(moviesChannel)
	for i, movie := range movies {
		moviesChannel <- fmt.Sprintf("%d. %s", i+1, movie)
	}
}
