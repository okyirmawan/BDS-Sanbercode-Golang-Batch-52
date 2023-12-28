package main

import (
	"fmt"
	"math"
	"sync"
)

type ShapeResult struct {
	JariJari float64
	Luas     float64
	Keliling float64
	Volume   float64
}

func calculateCircleProperties(jariJari float64, tinggi float64, resultChannel chan<- ShapeResult, wg *sync.WaitGroup) {
	defer wg.Done()

	luas := math.Pi * math.Pow(jariJari, 2)
	keliling := 2 * math.Pi * jariJari
	volume := luas * tinggi

	result := ShapeResult{
		JariJari: jariJari,
		Luas:     luas,
		Keliling: keliling,
		Volume:   volume,
	}

	resultChannel <- result
}

func main() {
	listJariJari := []float64{8, 14, 20}
	tinggi := 10.00

	resultChannel := make(chan ShapeResult)
	var wg sync.WaitGroup

	for _, jariJari := range listJariJari {
		wg.Add(1)
		go calculateCircleProperties(jariJari, tinggi, resultChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		fmt.Printf("Jari Jari: %.2f\n", result.JariJari)
		fmt.Printf("Luas: %.2f\n", result.Luas)
		fmt.Printf("Keliling: %.2f\n", result.Keliling)
		fmt.Printf("Volume: %.2f\n\n", result.Volume)
	}
}
