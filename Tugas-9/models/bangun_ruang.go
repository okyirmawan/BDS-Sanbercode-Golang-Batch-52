package models

import "math"

// Tabung
type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return (2 * math.Pi * t.JariJari) * (t.JariJari + t.Tinggi)
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * ((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi)))
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}
