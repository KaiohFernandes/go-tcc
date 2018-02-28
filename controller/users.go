package controller

import (
	"fmt"
	"net/http"
)

type User struct {}

func (user *User) Index(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Hello")
}

func (user *User) Sobre(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sobre")
}

func (user *User) Contato(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contato")
}