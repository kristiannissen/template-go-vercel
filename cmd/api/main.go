package main

import (
	"flag"
	"log"
	"net/http"
	api "template-go-vercel/api"
)

var port string

func init() {
	flag.StringVar(&port, "port", "Port", "The port man...")
	flag.Parse()
}

func main() {
	// New mux
	mux := http.NewServeMux()
	// Route
	mux.Handle("/api/hello", http.HandlerFunc(api.Hello))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
