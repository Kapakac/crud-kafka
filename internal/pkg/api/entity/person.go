package entity

import (
	"gorm.io/gorm"
)

type QuasiPersons struct {
	gorm.Model `json:"model"`
	Name       string `json:"Name"`
}
