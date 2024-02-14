package routes

import (
	"net/http"
	"strings"

)

/*  The Router can handle all the incomming requests.
*   For handling the request use router.HandleFunc. It takes the path and a handler function to manage the request
*
**/
type Router struct{
    opts RouterOpts
    path Path
}


func NewRouter() Router{
    return Router{
    path: EmptyPath(),
    }
}



// Interface for http
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
    var handler http.Handler
    //check if path is valid
    pathContent := router.path.Search(strings.Split(r.URL.Path, "/"))
    if(pathContent != nil){
        handler = pathContent.handler
    }else{
        handler=http.NotFoundHandler()
    }

    // for _,route := range router.routes{
    //   if(route.path == r.URL.Path){
    //     route.HandlerFunc(middlewares.Chain(route.handlerFunc,route.middlewares...))
    //   handler = route.Handler()
    // break
    //}
    //}

    handler.ServeHTTP(w,r)
}


// Handles a new path with given Handler
func (r *Router) HandleFunc(path string, handler http.HandlerFunc){
    routes := strings.Split(path, "/")
    r.path.AddPath(routes,handler)
}

// Options struct for the router
type RouterOpts struct{
    
    
}
