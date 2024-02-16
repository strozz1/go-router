package routes

import (
	"errors"
	"log"
	"net/http"
)

// The Router can handle all the incomming requests.
// For handling the request use router.HandleFunc. It takes the path and a handler function to manage the request
type Router struct {
	opts   RouterOpts
	routes Tree[Segment]
}

func NewRouter() Router {
	return Router{
		routes: EmptyTree[Segment](),
	}
}


// For debug purposes
func (r *Router) PrintRoutes() { 
	r.routes.Print()
}


// splits the given path and returns tuple with the slice in case of success or error.
//
// The returned slice has '/' on each element
//
// Output style: ['/home','/user','/posts']
func SplitPath(route string) ([]string, error) {
	routes := []string{}

	if len(route) == 0 {
		return nil, errors.New("route must not be empty")
	}
	if route[0] != '/' {
		return nil, errors.New("route must start with '/'")
	}

	segment := ""
	for _, c := range route {
		if c == '/' {
			if len(segment) > 1 {

				routes = append(routes, segment)
			}
			segment = "/"
		} else {
			segment = segment + string(c)
		}
	}
	routes = append(routes, segment)

	return routes, nil
}

// Interface for http
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	//check if path is valid
	path, err := SplitPath(r.URL.Path)
	if err != nil {
		handler = http.NotFoundHandler()
	} else {
		pathContent := router.Search(path)
		if pathContent != nil && pathContent.handler != nil{
			handler = pathContent.handler
		} else {
			handler = http.NotFoundHandler()
		}
	}

	// for _,route := range router.routes{
	//   if(route.path == r.URL.Path){
	//     route.HandlerFunc(middlewares.Chain(route.handlerFunc,route.middlewares...))
	//   handler = route.Handler()
	// break
	//}
	//}

	handler.ServeHTTP(w, r)
}

// Creates a new endpoint for the http router with the given route and the handler for the endpoint.
// The handler must be of type http.HandlerFunc
func (r *Router) Endpoint(path string, handler http.HandlerFunc) {
	paths, err := SplitPath(path)
	if err != nil {
		log.Fatal(err)
	}
	r.AddRoute(paths, handler)
}


// Adds a route to the current router. 
//
// @params: routes-[]string the route, handler: handler func for the route
func (r *Router) AddRoute(routes []string, handler http.HandlerFunc) *Segment {

	current := r.routes.root
	var tmp Segment
	for _, n := range routes {
		tmp = FromString(n)
		find := current.GetChild(tmp)

		//if nil create new node and append to current
		if find == nil {
			find = &Node[Segment]{
				value:    tmp,
				children: []*Node[Segment]{},
			}
			current.children = append(current.children, find)
		}

		current = find
	}
	content := &current.value
	content.handler = handler
	return content
}

// searchs for a Segment matching the specified route.
// if segment not found, nil is returned
func (r *Router) Search(routes []string) *Segment {
	current := r.routes.root
	var tmp Segment
	for _, n := range routes {
		tmp = FromString(n)
		find := current.GetChild(tmp)

		if find == nil {
			return nil
		}
		current = find
	}
	return &current.value
}

// Options struct for the router
type RouterOpts struct {
}
