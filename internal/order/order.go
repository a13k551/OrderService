package order

import (
	"encoding/json"
)

type Order struct {
	Id  string  `json:"id"`
	Num string  `json:"num"`
	Sum float32 `json:"sum"`
}

func NewOrder(id, num string, sum float32) *Order {
	return &Order{
		Id:  id,
		Num: num,
		Sum: sum,
	}
}

func GetOrder() ([]byte, error) {

	order1 := NewOrder("1", "1", 100.25)

	js, err := json.Marshal(order1)

	return js, err

}
