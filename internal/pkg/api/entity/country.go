package entity

import (
	"github.com/Kapakac/crud-kafka/internal/pkg/api/query"
	"gorm.io/gorm"
)

type QuasiCountry struct {
	gorm.Model
	Country string `json:"country"`
}

func GetCountries() interface{} {
	return query.Finder{Arr: []QuasiCountry{}}.All()
}
