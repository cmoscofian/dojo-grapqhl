// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Claim struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Stage      string `json:"stage"`
	Status     string `json:"status"`
	ClientID   int    `json:"client_id"`
	ResourceID int    `json:"resource_id"`
	Resource   string `json:"resource"`
	ReasonID   string `json:"reason_id"`
	Fulfilled  bool   `json:"fulfilled"`
	Order      *Order `json:"order"`
}

type Item struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	CategoryID string `json:"category_id"`
}

type Order struct {
	ID         int            `json:"id"`
	Shipping   *OrderShipping `json:"shipping"`
	Shipment   *Shipment      `json:"shipment"`
	OrderItems []*OrderItem   `json:"order_items"`
}

type OrderItem struct {
	Item       *Item   `json:"item"`
	Quantity   float64 `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	CurrencyID string  `json:"currency_id"`
}

type OrderShipping struct {
	ID int `json:"id"`
}

type Shipment struct {
	ID        int     `json:"id"`
	Mode      string  `json:"mode"`
	Status    string  `json:"status"`
	Substatus *string `json:"substatus"`
}
