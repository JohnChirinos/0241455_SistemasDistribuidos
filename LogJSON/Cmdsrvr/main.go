package main

import (
	"log"

	"LogJSON/Internalsrvr/server"
)

func main() {
	srv := server.NewHTTPServer(":3333")
	log.Fatal(srv.ListenAndServe())
}
