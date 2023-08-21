package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a13k551/OrderService/internal/order"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", simpleAnswer)
	log.Fatal(http.ListenAndServe("localhost:5000", mux))

}

func simpleAnswer(w http.ResponseWriter, req *http.Request) {

	order1 := order.NewOrder("1", "1", 100.25)

	js, err := json.Marshal(order1)

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
