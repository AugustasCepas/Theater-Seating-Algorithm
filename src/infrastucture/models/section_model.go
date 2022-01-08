package models

type Section struct {
	Id        int    `json:"id" validate:"gte=1"`
	LayoutId  int    `json:"layoutId" validate:"gte=1"`
	Name      string `json:"name" validate:"len=3"`
	HighSeats bool   `json:"highSeats"` // (e.g. on balcony)
	Curved    bool   `json:"curved"`
}
