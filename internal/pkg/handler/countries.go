package handler

import (
	"github.com/Kapakac/crud-kafka/internal/pkg/api/entity"
	"github.com/Kapakac/crud-kafka/internal/pkg/brokerkafka/producer"
	"github.com/Kapakac/crud-kafka/internal/pkg/database"
	"github.com/Kapakac/crud-kafka/internal/pkg/helper"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries := entity.GetCountries()

	producer.WriteMessages(map[string]interface{}{"test-get": countries})

	helper.RespondJSON(w, http.StatusOK, countries)
}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	country := getCountryOr404(uint(id), w, r)
	if country == nil {
		return
	}

	helper.RespondJSON(w, http.StatusOK, country)
}

func getCountryOr404(id uint, w http.ResponseWriter, r *http.Request) *entity.QuasiCountry {
	country := entity.QuasiCountry{}
	if err := database.DB.First(&country, entity.QuasiCountry{Model: gorm.Model{ID: id}}).Error; err != nil {
		helper.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &country
}
