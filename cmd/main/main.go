package main

import (
	"net"
	"net/http"
	"time"

	"github.com/a13k551/OrderService/internal/order"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.StrictSlash(true)

	order.AddHandlers(router)

	listener, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		panic(err)
	}

	serv := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = serv.Serve(listener)

	if err != nil {
		panic(err)
	}

}
