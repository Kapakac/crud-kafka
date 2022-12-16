package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Get(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(path, f).Methods("GET")
}

func Post(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(path, f).Methods("POST")
}
