package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/a13k551/OrderService/internal/order"
	"github.com/gorilla/mux"
)

func OrderHandler(w http.ResponseWriter, req *http.Request) {

}

func GetOrderById(w http.ResponseWriter, req *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(req)["id"])

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

func CreateOrder(w http.ResponseWriter, req *http.Request) {

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

func DeleteOrderById(w http.ResponseWriter, req *http.Request) {

}

func UpdateOrderById(w http.ResponseWriter, req *http.Request) {

}
