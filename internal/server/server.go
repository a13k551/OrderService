package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/a13k551/OrderService/internal/order"
)

func OrderHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/order/" {

		if req.Method == http.MethodPost {
			createOrder(w, req)
		} else if req.Method == http.MethodGet {
			getOrder(w, req)
		} else if req.Method == http.MethodDelete {
			deleteOrder(w, req)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method),
				http.StatusMethodNotAllowed)
			return
		}
	}
}

func getOrder(w http.ResponseWriter, req *http.Request) {

	js, err := order.GetOrder()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func createOrder(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	ord := &order.Order{}

	err = json.Unmarshal(body, ord)

	if err != nil {
		fmt.Println(err)
	}
}

func deleteOrder(w http.ResponseWriter, req *http.Request) {
}