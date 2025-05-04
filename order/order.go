package order

type Order struct {
	Id        int     `json:"id"`
	UserId    float64 `json:"userId"`
	ItemName  string  `json:"itemName"`
	Amount    float32 `json:"amount"`
	CreatedAt int     `json:"createdAt"`
}
