package models

type Row struct {
	Id        int `json:"id" validate:"gte=1"`
	SectionId int `json:"section" validate:"gte=1"`
	Number    int `json:"number" validate:"gte=1"`
}
