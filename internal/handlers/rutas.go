package handlers

import (
	"fmt"
	"net/http"
)

func HandleRutas(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Rutas...")
}

func HandleRutasUnit(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Ruta: %+v\n", r.URL.Path)
}
