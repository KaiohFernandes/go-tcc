package routes

import (
	"fmt"
	"net/http"
)

func index(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Hello")
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Teste")
}
