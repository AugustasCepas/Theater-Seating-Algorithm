package entities

type Input struct {
	Reserve string `json:"reserve" validate:"gte=1"`
}
