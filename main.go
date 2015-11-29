package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":5000"
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(port, nil)
	log.Println("Servidor de TODO corriendo en localhost:%s", port)
}
