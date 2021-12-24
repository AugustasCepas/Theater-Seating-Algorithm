package models

type Section struct {
	Id        int64  `json:"id" validate:"gte=1"`
	LayoutId  int64  `json:"layoutId" validate:"gte=1"`
	Name      string `json:"name" validate:"len=3"`
	HighSeats bool   `json:"highSeats"`
	Curved    bool   `json:"curved"`
}
