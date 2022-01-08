package models

import (
	"github.com/shopspring/decimal"
)

type Rank struct {
	Id    int             `json:"id" validate:"gte=1"`
	Price decimal.Decimal `json:"price"`
	Row   int             `json:"row" validate:"gte=1"`
}
