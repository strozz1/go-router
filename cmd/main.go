package main

import (
	"log"
	"net/http"

	"github.com/strozz1/go-router/internal/handlers"
	"github.com/strozz1/go-router/internal/routes"
)

func main() {

    //Example usage
    addr := ":8000"   
  	router := routes.NewRouter()

	router.Endpoint("/", handlers.HandleIndex)
    router.Endpoint("/inicio", handlers.HandleHome)      //.Middlewares(middlewares.Method("GET"),middlewares.Logging())
	router.Endpoint("/data", handlers.HandleRutas)      //.Middlewares(middlewares.Method("GET"),middlewares.Logging())
	router.Endpoint("/user/pendientes", handlers.HandleRutas) //.Middlewares(middlewares.Method("GET"),middlewares.Logging())
	router.Endpoint("/test", handlers.HandleRutasUnit) //.Middlewares(middlewares.Method("GET"),middlewares.Logging())
	router.Endpoint("/user/{id}", handlers.HandleRutasUnit) //.Middlewares(middlewares.Method("GET"),middlewares.Logging())

    router.PrintRoutes()   
    log.Println("Listening on ",addr)
	http.ListenAndServe(addr, &router)
}
