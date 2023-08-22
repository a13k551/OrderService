package main

import (
	"log"
	"net/http"

	"github.com/a13k551/OrderService/internal/server"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/order/", server.OrderHandler)
	log.Fatal(http.ListenAndServe("localhost:5000", mux))
}
