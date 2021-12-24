package models

type Row struct {
	Id      int64  `json:"id" validate:"gte=1"`
	Name    string `json:"name" validate:"len=3"`
	Section int64  `json:"section" validate:"gte=1"`
	Rank    int64  `json:"rank" validate:"gte=1"`
	Seats   []int  `json:"seats"`
}
