package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-14/config"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-14/models"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-14/nilai_mahasiswa"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-14/utils"
	"log"
	"net/http"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()
	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Read
// GetNilai
func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilai, err := nilai_mahasiswa.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
}

// Create
// PostNilai
func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilai models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	nilai.IndeksNilai, _ = getIndexNilai(nilai.Nilai)
	if err := nilai_mahasiswa.Insert(ctx, nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Update
// UpdateNilai
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilai models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	nilai.IndeksNilai, _ = getIndexNilai(nilai.Nilai)
	if err := nilai_mahasiswa.Update(ctx, nilai, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteNilai
func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idNilai = ps.ByName("id")
	if err := nilai_mahasiswa.Delete(ctx, idNilai); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
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
