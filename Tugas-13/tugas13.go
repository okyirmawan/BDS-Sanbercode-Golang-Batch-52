package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	apiUsername = "admin"
	apiPassword = "password"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var listNilaiMahasiswa = make([]NilaiMahasiswa, 0)

func authenticate(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		return false
	}

	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		return false
	}

	username := credentials[0]
	password := credentials[1]

	return username == apiUsername && password == apiPassword
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Authenticate the request
	if !authenticate(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var newNilai NilaiMahasiswa
	var err error

	if r.Header.Get("Content-Type") == "application/json" {
		// parse dari json
		decodeJSON := json.NewDecoder(r.Body)
		if err := decodeJSON.Decode(&newNilai); err != nil {
			log.Fatal(err)
		}
	} else {
		// parse dari form
		getNilai := r.PostFormValue("Nilai")
		nilai, _ := strconv.Atoi(getNilai)
		newNilai.Nilai = uint(nilai)
		newNilai.Nama = r.PostFormValue("Nama")
		newNilai.MataKuliah = r.PostFormValue("MataKuliah")
	}

	newNilai.IndeksNilai, err = getIndexNilai(newNilai.Nilai)
	newNilai.ID = uint(len(listNilaiMahasiswa) + 1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listNilaiMahasiswa = append(listNilaiMahasiswa, newNilai)

	response, err := json.Marshal(newNilai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func getIndexNilai(nilai uint) (string, error) {
	switch {
	case nilai > 100:
		return "", errors.New("Maaf nilai tidak boleh melebihi 100")
	case nilai >= 80:
		return "A", nil
	case nilai >= 70:
		return "B", nil
	case nilai >= 60:
		return "C", nil
	case nilai >= 50:
		return "D", nil
	default:
		return "E", nil
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(listNilaiMahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func dataNilai(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePostRequest(w, r)
	case "GET":
		handleGetRequest(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/nilai-mahasiswa", dataNilai)
	fmt.Println("Server running at http://localhost:8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
