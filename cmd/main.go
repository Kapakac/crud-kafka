package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := App{}
	a.Initialize()
	a.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test", r.URL.Path)
	})
	a.Run(":8064")

}
