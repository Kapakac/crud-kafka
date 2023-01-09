package entity

import (
	"gorm.io/gorm"
)

type QuasiPersons struct {
	gorm.Model
	Name string `json:"Name"`
}
