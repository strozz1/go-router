package api

import (
	"net/http"

	"github.com/strozz1/pinkbikers-web/internal/handlers"
	"github.com/strozz1/pinkbikers-web/internal/routes"
	"github.com/strozz1/pinkbikers-web/internal/storage"
	"github.com/strozz1/pinkbikers-web/internal/types"
)


type Server struct{
    addr string
    storage storage.Storage[types.Ruta]

}

func New(addr string,stor storage.Storage[types.Ruta]) *Server{
    return &Server{
        addr: addr,
        storage: stor,
    }
}


func (s *Server) Start() error{

    router := routes.NewRouter()
    
    router.HandleFunc("/",handlers.HandleIndex)
    router.HandleFunc("/inicio",handlers.HandleHome)//.Middlewares(middlewares.Method("GET"),middlewares.Logging())
    router.HandleFunc("/rutas",handlers.HandleRutas)//.Middlewares(middlewares.Method("GET"),middlewares.Logging())
    router.HandleFunc("/rutas/{id}",handlers.HandleRutas)//.Middlewares(middlewares.Method("GET"),middlewares.Logging())

    
    http.ListenAndServe(s.addr,&router)
    return nil
}


