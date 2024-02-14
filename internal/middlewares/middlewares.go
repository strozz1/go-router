package middlewares

import (
	"log"
	"net/http"
)



type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, midls ...Middleware) http.HandlerFunc{
    for _, m:=range midls{
        f=m(f)
    }
    
    return f

}


func Logging() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc{
        
        return func(w http.ResponseWriter, r *http.Request){
            log.Println(r.URL.Path)
            f(w,r)
        }
    }

}



// Checks if http request matches the method provided.
func Method(m string) Middleware{
    
        return func(f http.HandlerFunc) http.HandlerFunc{

            return func(w http.ResponseWriter, r *http.Request){

                if(r.Method != m){
                    http.Error(w,http.StatusText(http.StatusBadRequest),http.StatusBadRequest)
                    return
                }

                f(w,r)
            }
        }
}
