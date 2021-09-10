package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type credential interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewRouter(c credential) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/credentials", c.Create).Methods(http.MethodPost)

	return r
}
