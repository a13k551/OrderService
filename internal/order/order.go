package order

import (
	"encoding/json"
)


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
