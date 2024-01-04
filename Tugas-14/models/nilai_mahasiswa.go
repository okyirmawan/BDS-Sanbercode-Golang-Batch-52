package models

import "time"

type (
	NilaiMahasiswa struct {
		ID          uint      `json:"id"`
		Nama        string    `json:"nama"`
		MataKuliah  string    `json:"mata_kuliah"`
		IndeksNilai string    `json:"indeks_nilai"`
		Nilai       uint      `json:"nilai"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
