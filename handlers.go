package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

func registrarHandler(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	ciudad := r.FormValue("ciudad")
	genero := r.FormValue("genero")
	if nombre == "" || ciudad == "" || genero == "" {
		log.Println("Campos sin valor encontrados")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	persona := &Persona{nombre, ciudad, genero}
	_, err := registrarPersona(persona)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Registro agregado")
}
func listarPersonasHandler(w http.ResponseWriter, _ *http.Request) {
	personas, err := personas()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(personas); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
