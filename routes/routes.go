package routes

import (
	"strings"
	"fmt"
	"net/http"
)

type Routes struct {}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Teste")
}

func (routes *Routes) RouteHandler(w http.ResponseWriter, r *http.Request){

	path := strings.TrimPrefix(r.URL.Path, "/")
	
	switch {
	case path == "":
		index(w, r)
		break
	case path == "teste":
		teste(w, r)
		break
	}

}