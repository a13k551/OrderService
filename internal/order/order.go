package order

type Order struct {
	id  string
	num string
	sum float32
}

func (*Order) New(id, num string, sum float32) *Order {
	return &Order{
		id:  id,
		num: num,
		sum: sum,
	}
}
