package model

import "github.com/jinzhu/gorm"

type Mahasiswa struct {
	gorm.Model
	NIM         string `json:"nim" gorm:"column:nim"`
	Nama        string `json:"nama" gorm:"column:nama"`
	TempatLahir string `json:"tempat_lahir" gorm:"column:tempat_lahir"`
	Agama       string `json:"agama" gorm:"column:agama"`
}
