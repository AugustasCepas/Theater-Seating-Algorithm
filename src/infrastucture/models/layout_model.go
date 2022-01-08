package models

type Layout struct {
	Id   int    `json:"id" validate:"gte=1"`
	Name string `json:"name" validate:"len=3"`
}
