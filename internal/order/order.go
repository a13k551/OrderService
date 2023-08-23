package order

import (
	"encoding/json"
)

type Order struct {
	Id  int     `json:"id"`
	Num string  `json:"num"`
	Sum float32 `json:"sum"`
}

func NewOrder(id int, num string, sum float32) *Order {
	return &Order{
		Id:  id,
		Num: num,
		Sum: sum,
	}
}

func GetOrderById(id int) ([]byte, error) {

	order1 := NewOrder(id, "1", 100.25)

	js, err := json.Marshal(order1)

	return js, err

}
