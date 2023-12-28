package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var jariJari float64 = 7
		var tinggi float64 = 10

		volume := volumeTabung(jariJari, tinggi)
		luas := luasLingkaran(jariJari)
		keliling := kelilingLingkaran(jariJari)

		text := fmt.Sprintf(" jariJari : %.0f, tinggi: %.0f, volume : %.2f, luas alas: %.2f, keliling alas: %.2f", jariJari, tinggi, volume, luas, keliling)
		_, err := fmt.Fprintln(w, text)
		if err != nil {
			return
		}
	})

	fmt.Println("starting web server at http://127.0.0.1:8888/")
	http.ListenAndServe(":8888", nil)
}

func luasLingkaran(jariJari float64) float64 {
	return math.Pi * math.Pow(jariJari, 2)
}

func kelilingLingkaran(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func volumeTabung(jariJari float64, tinggi float64) float64 {
	return luasLingkaran(jariJari) * tinggi
}
