package main

import (
	"net/http";
	"os"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)

type Persona struct {
	Nombre string `json:"nombre,omitempty"`
	Ciudad string `json:"ciudad,omitempty"`
	Genero string `json:"genero,omitempty"`
}

func main() {
	conectarBaseDatos()
	defer cerrarBaseDatos()

	router := mux.NewRouter()
	router.HandleFunc("/registrar", registrarHandler).Methods("POST")
	router.HandleFunc("/personas", listarPersonasHandler).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	http.Handle("/",router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
