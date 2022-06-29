package entity

import (
	"gorm.io/gorm"
)

type QuasiCountries struct {
	gorm.Model `json:"model"`
	Country    string `json:"country"`
}
