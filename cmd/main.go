package main

import (
	"context"
	"door-greeter/scan_service/data"
	"door-greeter/scan_service/server"
	"log"
	"net/http"
)

func main() {
	data.DatabaseInit()

	ctx := context.Background()
	mux := &http.ServeMux{}
	server.SetRoutes(mux)
	srv := server.NewServer(mux)

	log.Fatal(server.StartServer(ctx, srv))
}
