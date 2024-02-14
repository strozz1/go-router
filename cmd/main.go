package main

import (
	"flag"
	"log"

	"github.com/strozz1/pinkbikers-web/internal/api"
	"github.com/strozz1/pinkbikers-web/internal/storage"
	"github.com/strozz1/pinkbikers-web/internal/types"
)

func main() {
    
    
    addr := flag.String("addr",":8000","The server address")
    flag.Parse()
    mock := &storage.MockDB[types.Ruta]{}
    server :=api.New(*addr,mock)
    log.Print("Server initialized at address: ", *addr)
    server.Start()


}
