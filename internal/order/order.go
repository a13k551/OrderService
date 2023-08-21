package order

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
