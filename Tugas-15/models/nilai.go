package models

type (
	Nilai struct {
		ID           uint   `json:"id"`
		Indeks       string `json:"indeks"`
		Skor         uint   `json:"skor"`
		MahasiswaID  uint   `json:"mahasiswa_id"`
		MataKuliahID uint   `json:"mata_kuliah_id"`
	}
)
