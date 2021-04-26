package main

import (
	"log"

	"github.com/mahdi/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":9999")
	log.Fatal(srv.ListenAndServe())
}
