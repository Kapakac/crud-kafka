package handler

import (
	"github.com/Kapakac/crud-kafka/internal/pkg/api/entity"
	"github.com/Kapakac/crud-kafka/internal/pkg/helper"
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries := entity.GetCountries()
	helper.RespondJSON(w, http.StatusOK, countries)
}
