package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama" validate:"required"`
	Email        string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required" format:"2006-01-02"`
	NoHP         string `json:"no_hp" form:"no_hp" validate:"required"`
	UrlFoto      string `json:"url_foto" form:"url_foto" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" gorm:"type:enum('Laki-laki', 'Perempuan');not_null" validate:"required,min=9"`
}
