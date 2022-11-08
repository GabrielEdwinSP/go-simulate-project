package domain

import "time"

type Request struct {
	ID        int       `json:"id" gorm:"primary_key"`
	GajiPokok float32   `json:"gajipokok"`
	Tunjangan float32   `json:"tunjangan"`
	CreatedAt time.Time `json:"createdat"`
}

type Response struct {
	ID           int       `json:"id" gorm:"primary_key"`
	TotalGross   float32   `json:"totalgross"`
	JHT          float32   `json:"jht"`
	JP           float32   `json:"jp"`
	BiayaJabatan float32   `json:"biayajabatan"`
	TotalNet     float32   `json:"totalnet"`
	GajiSetahun  float32   `json:"gajisetahun"`
	TotalSetahun float32   `json:"totalSetahun"`
	PPHperTahun  float32   `json:"pphpertahun"`
	PPHperbulan  float32   `json:"pphperbulan"`
	CreatedAt    time.Time `json:"createdat"`
}
