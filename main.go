package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/DeanCash/fl0_hello_api/routes"
)

var (
	host    string = "127.0.0.1"
	port    int    = 5555
	address string = net.JoinHostPort(host, fmt.Sprintf("%d", port))
)

func init() {}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	server.HandleFunc("/healthcheck", routes.HealthCheckHandler)

	// .HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)

	// 	w.Write([]byte("Hello, World!"))
	// })

	log.Println("-- Starting server")
	if err := http.ListenAndServe(address, server); err != nil {
		log.Fatalln("-- Error starting server:", err)
	}
}
