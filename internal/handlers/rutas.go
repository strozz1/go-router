package handlers

import (
    "fmt"
    "net/http"
    "github.com/strozz1/pinkbikers-web/internal/types"
)

func HandleRutas(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w,"Rutas...")
}

func HandleRutasUnit(w http.ResponseWriter, r *http.Request){

    ruta:= types.New("2345","Hola")
    fmt.Fprintf(w,"Ruta: %+v\n",ruta)
}
