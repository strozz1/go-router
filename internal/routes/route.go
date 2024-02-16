package routes

import (
	"net/http"

	"github.com/strozz1/go-router/internal/middlewares"
)

type Route struct{
    path string
    middlewares []middlewares.Middleware
    handlerFunc http.HandlerFunc
    handler http.Handler
    routerOpts RouterOpts

}
// Sets the path of the route
func (r *Route) Path(path string) *Route{
    r.path=path
    return r
}

// Add middlewares to the route
func (r *Route) Middlewares(middlewares ...middlewares.Middleware){
    for _,m := range middlewares{
        r.middlewares=append(r.middlewares, m)
    }
}

func (r *Route) Handler() http.Handler{
    return http.Handler(r.handlerFunc)
}


//sets to the hanlder the hanlder function
func (r *Route) HandlerFunc(f http.HandlerFunc) *Route{
        r.handlerFunc= f
        return r
}

