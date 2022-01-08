package usecases

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AugustasCepas/Theater-Seating-Algorithm/domain/repositories"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/dtos"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/infrastucture/models"
)

type SeatUsecases struct {
	SeatRepository repositories.SeatRepository
	Seat           models.Seat
}

func NewSeatUsecases(seatRepository repositories.SeatRepository) *SeatUsecases {
	return &SeatUsecases{SeatRepository: seatRepository}
}

func (ac SeatUsecases) ReserveSeats(layoutId int, sectionId int, rankId int, input string) (int, error) {
	totalSeatsReserved := 0

	reservationsSlice, err := getReservationsSlice(input)
	if err != nil {
		return totalSeatsReserved, err
	}

	lastReservationId, err := ac.SeatRepository.GetLastReservationId()
	if err != nil {
		return totalSeatsReserved, err
	}

	sectionSeats, err := ac.SeatRepository.GetSectionSeats(layoutId, sectionId, rankId)
	if err != nil {
		return totalSeatsReserved, err
	}

	err = areEnoughSeatsInTheSection(reservationsSlice, sectionSeats)
	if err != nil {
		return totalSeatsReserved, err
	}

	var oddRowSeats []dtos.SectionSeats
	for _, users := range reservationsSlice {

		for i := 0; i < users; i++ {
			var seatToReserveId int

			if sectionSeats[totalSeatsReserved].RowId%2 == 1 {
				seatToReserveId = sectionSeats[totalSeatsReserved].Id
			} else {
				if len(oddRowSeats) == 0 {
					seat := totalSeatsReserved
					for sectionSeats[seat].RowId%2 == 0 {
						oddRowSeats = append(oddRowSeats, sectionSeats[seat])
						if seat >= len(sectionSeats)-1 {
							break
						}
						seat++
					}
				}

				seatToReserveId = oddRowSeats[len(oddRowSeats)-1].Id
				oddRowSeats = oddRowSeats[:len(oddRowSeats)-1]
			}

			rowsAffected, err := ac.SeatRepository.ReserveSeat(lastReservationId+1, seatToReserveId)
			if err != nil || rowsAffected == 0 {
				return int(rowsAffected), err

			}
			totalSeatsReserved++
		}
		lastReservationId++
	}

	return totalSeatsReserved, nil
}

func areEnoughSeatsInTheSection(reservationsSlice []int, sectionSeats []dtos.SectionSeats) error {
	var usersToSit int
	for _, users := range reservationsSlice {
		usersToSit += users
	}

	if len(sectionSeats) < usersToSit {
		return errors.New("error, there are not enough seats in the section")
	}

	return nil
}

func (ac SeatUsecases) GetLayoutSeats(layoutId int, sectionId int) ([]string, error) {
	seats, err := ac.SeatRepository.GetLayoutSeats(layoutId, sectionId)
	if err != nil {
		return nil, err
	}

	layoutRows, err := ac.SeatRepository.GetLayoutRows(layoutId, sectionId)
	if err != nil {
		return nil, err
	}

	seatsInARow := len(seats) / layoutRows

	layout := CreateLayout(seatsInARow, layoutRows)
	layout, err = FillLayout(layout, seats, seatsInARow)
	if err != nil {
		return nil, err
	}

	response := CreateResponse(layout)

	return response, nil
}

func CreateResponse(layout [][]int) []string {
	var response string
	var responeSlice []string
	for i := 0; i < 3; i++ {
		for j := 0; j < 8; j++ {
			response += strconv.Itoa(layout[i][j])
		}
		responeSlice = append(responeSlice, response)
		response = ""
	}
	return responeSlice
}

func CreateLayout(seatsInARow int, rows int) [][]int {

	layout := make([][]int, rows)
	for i := 0; i < rows; i++ {
		layout[i] = make([]int, seatsInARow)
	}

	return layout
}

func FillLayout(layout [][]int, seats []dtos.LayoutSeats, seatsInARow int) ([][]int, error) {
	currentRow := 0
	currentSeat := 0

	for i, _ := range seats {

		layout[currentRow][currentSeat] = seats[i].ReservationId

		if currentSeat == seatsInARow-1 {
			currentSeat = 0
			currentRow++
		} else {
			currentSeat++
		}
	}

	return layout, nil
}

func (ac SeatUsecases) GetUserSeats(reservationId int) ([]dtos.ReservedSeats, error) {
	seats, err := ac.SeatRepository.GetReservationSeats(reservationId)
	if err != nil {
		return seats, err
	}

	return seats, nil
}

func getReservationsSlice(input string) ([]int, error) {
	sInput := strings.Split(input, ",")
	iInput := make([]int, len(sInput))

	for i, v := range sInput {

		inputSlice, err := strconv.Atoi(v)
		if inputSlice == 0 || err != nil {
			return iInput, errors.New("error: reserve input contains invalid value")
		}
		iInput[i], _ = inputSlice, err
	}

	return iInput, nil
}
