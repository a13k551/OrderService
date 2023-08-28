package order

type Order struct {
	Id  int     `json:"id"`
	Num string  `json:"num"`
	Sum float32 `json:"sum"`
}
