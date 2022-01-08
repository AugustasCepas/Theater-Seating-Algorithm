package repositories

import (
	"github.com/AugustasCepas/Theater-Seating-Algorithm/dtos"
)

type SeatRepository interface {
	ReserveSeat(reservationId int, seatId int) (int64, error)
	GetReservationSeats(reservationId int) ([]dtos.ReservedSeats, error)
	GetLayoutSeats(layoutId int, sectionId int) ([]dtos.LayoutSeats, error)

	GetSectionSeats(layoutId int, sectionId int, rankId int) ([]dtos.SectionSeats, error)
	GetLastReservationId() (int, error)
	GetLayoutRows(layoutId int, sectionId int) (int, error)
}
