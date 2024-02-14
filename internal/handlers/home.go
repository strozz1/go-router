package handlers

import (
    "net/http"
    "fmt"
)


func HandleHome(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w,"Home page")
}

func HandleIndex(w http.ResponseWriter, r *http.Request){
    http.Redirect(w,r,"/inicio",http.StatusPermanentRedirect)

    
}

