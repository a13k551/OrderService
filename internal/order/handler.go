package order

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddHandlers(router *mux.Router){
	router.HandleFunc("/order/", handlerCreate).Methods("POST")
	router.HandleFunc("/order/{id:[0-9]+}/", handlerGetById).Methods("GET")
	router.HandleFunc("/order/{id:[0-9]+}/", handlerDeleteById).Methods("DELETE")
	router.HandleFunc("/order/{id:[0-9]+}/", handlerUpdateById).Methods("UPDATE")
}

func handlerCreate(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	ord := &Order{}

	err = json.Unmarshal(body, ord)

	if err != nil {
		fmt.Println(err)
	}
}

func handlerGetById(w http.ResponseWriter, req *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	js, err := GetOrderById(id)

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



func handlerDeleteById(w http.ResponseWriter, req *http.Request) {
}

func handlerUpdateById(w http.ResponseWriter, req *http.Request) {
}
