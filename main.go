package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/DeanCash/fl0_hello_api/routes"
)

const (
	PORT_ENV string = "PORT"
)

var (
	host    string
	port    string
	address string
)

func init() {
	_port, err := os.LookupEnv(PORT_ENV)
	if !err {
		log.Panicln("There is no PORT environment variable!")
	}

	host = "0.0.0.0"
	port = _port
	address = net.JoinHostPort(host, port)
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	server.HandleFunc("/healthcheck", routes.HealthCheckHandler)

	log.Println("-- Starting server")
	if err := http.ListenAndServe(address, server); err != nil {
		log.Fatalln("-- Error starting server:", err)
	}
}
