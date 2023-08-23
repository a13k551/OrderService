package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/a13k551/OrderService/internal/order"
)

func OrderHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/order/" {

		if req.Method == http.MethodPost {
			createOrder(w, req)
		}
	} else {

		path := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /task/<id> in task handler", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Method == http.MethodDelete {
			deleteOrderById(w, req, int(id))
		} else if req.Method == http.MethodGet {
			getOrderById(w, req, int(id))
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /task/<id>, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func getOrderById(w http.ResponseWriter, req *http.Request, id int) {

	js, err := order.GetOrderById(id)

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

func deleteOrderById(w http.ResponseWriter, req *http.Request, id int) {

}
