package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/a13k551/OrderService/internal/config"
	"github.com/a13k551/OrderService/internal/order"
	"github.com/gorilla/mux"
)

func main() {

	conf := config.Get()

	router := mux.NewRouter()
	router.StrictSlash(true)

	order.AddHandlers(router)

	address := fmt.Sprintf("%s:%s", conf.App.BindIP, conf.App.Port)
	listener, err := net.Listen("tcp", address)

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
