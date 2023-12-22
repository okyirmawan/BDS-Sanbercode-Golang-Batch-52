package models

// Segitiga Sama Sisi
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

// Persegi Panjang
type PersegiPanjang struct {
	Panjang, Lebar int
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int {
	return 2 * (pp.Panjang + pp.Lebar)
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}
