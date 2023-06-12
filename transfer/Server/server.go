package Server

import (
	s "GoMango/src/web"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.Index).Methods("GET")
	r.HandleFunc("/", s.Submit).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	http.ListenAndServe(":8080", r)
}
