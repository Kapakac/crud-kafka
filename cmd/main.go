package main

import (
	"github.com/Kapakac/crud-kafka/internal/pkg/entity"
	"github.com/Kapakac/crud-kafka/internal/pkg/entity/query"
)

func main() {
	a := App{}
	a.Initialize()

	query.Finder{Arr: []entity.QuasiCountries{}}.All()
	//a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "test", r.URL.Path)
	//})
	a.Run("127.0.0.1:8064")
}
