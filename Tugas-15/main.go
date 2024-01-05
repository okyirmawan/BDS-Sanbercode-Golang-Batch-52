package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/config"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/mahasiswa"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/mata_kuliah"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/models"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/nilai"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/utils"
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

	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa/create", PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", DeleteMahasiswa)

	router.GET("/matakuliah", GetMataKuliah)
	router.POST("/matakuliah/create", PostMataKuliah)
	router.PUT("/matakuliah/:id/update", UpdateMataKuliah)
	router.DELETE("/matakuliah/:id/delete", DeleteMataKuliah)

	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/**
Nilai
*/

// Read
// GetNilai
func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilai, err := nilai.GetAll(ctx)

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

	var mNilai models.Nilai
	var err error

	if err := json.NewDecoder(r.Body).Decode(&mNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	mNilai.Indeks, err = getIndeksNilai(mNilai.Skor)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := nilai.Insert(ctx, mNilai); err != nil {
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
	var mNilai models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&mNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")
	var err error

	mNilai.Indeks, err = getIndeksNilai(mNilai.Skor)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := nilai.Update(ctx, mNilai, idNilai); err != nil {
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
	if err := nilai.Delete(ctx, idNilai); err != nil {
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

/**
Mahasiswa
*/

// Read
// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mhs, err := mahasiswa.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, mhs, http.StatusOK)
}

// Create
// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := mahasiswa.Insert(ctx, mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Update
// UpdateMahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMahasiswa = ps.ByName("id")

	if err := mahasiswa.Update(ctx, mhs, idMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMahasiswa
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMahasiswa = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idMahasiswa); err != nil {
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

/**
Mata Kuliah
*/

// Read
// GetMataKuliah
func GetMataKuliah(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	matkul, err := mata_kuliah.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, matkul, http.StatusOK)
}

// Create
// PostMataKuliah
func PostMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := mata_kuliah.Insert(ctx, mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Update
// UpdateMataKuliah
func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMataKuliah = ps.ByName("id")

	if err := mata_kuliah.Update(ctx, mhs, idMataKuliah); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMataKuliah
func DeleteMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMataKuliah = ps.ByName("id")
	if err := mata_kuliah.Delete(ctx, idMataKuliah); err != nil {
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

func getIndeksNilai(nilai uint) (string, error) {
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
