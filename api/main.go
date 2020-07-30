package main

import (
	"cbm-api/controllers"
	"log"
	"net/http"
)

//Start Server -> Serve routes -> Defer server destroy
func main() {
	// TODO: add go routine for watcher.
	// Find how to request to api
	server, closer := controllers.NewServer()

	defer closer()
	log.Println("Server runs on http://localhost:" + server.Port)
	log.Print(
		http.ListenAndServe(":"+server.Port, server.HandelerCores()(server.Router)))
}