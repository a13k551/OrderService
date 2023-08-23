package main

import (
	"net"
	"net/http"
	"time"

	"github.com/a13k551/OrderService/internal/server"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/order/", server.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{id:[0-9]+}/", server.GetOrderById).Methods("GET")
	router.HandleFunc("/order/{id:[0-9]+}/", server.DeleteOrderById).Methods("DELETE")
	router.HandleFunc("/order/{id:[0-9]+}/", server.UpdateOrderById).Methods("UPDATE")

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
