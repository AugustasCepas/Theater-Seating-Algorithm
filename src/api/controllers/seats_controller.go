package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AugustasCepas/Theater-Seating-Algorithm/domain/repositories/entities"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/usecases"
	"github.com/gorilla/mux"
)

type SeatsController struct {
	SeatUsecases usecases.SeatUsecases
}

func NewSeatController(seatUsecases usecases.SeatUsecases) *SeatsController {
	return &SeatsController{SeatUsecases: seatUsecases}
}

func GetReservationData(vars map[string]string) (int, int, int, error) {
	layoutId, err := strconv.Atoi(vars["layoutId"])
	if err != nil {
		return 0, 0, 0, err
	}

	sectionId, err := strconv.Atoi(vars["sectionId"])
	if err != nil {
		return 0, 0, 0, err
	}

	rankId, err := strconv.Atoi(vars["rankId"])
	if err != nil {
		return 0, 0, 0, err
	}

	return layoutId, sectionId, rankId, err
}

func (sc SeatsController) ReserveSeats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	layoutId, sectionId, rankId, err := GetReservationData(vars)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var input entities.Input
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	rowsAffected, err := sc.SeatUsecases.ReserveSeats(layoutId, sectionId, rankId, input.Reserve)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if rowsAffected == 0 {
		respondWithError(w, http.StatusBadRequest, "Seat(s) not reserved")
		return
	}

	response := fmt.Sprintf("%d Seat(s) reserved", rowsAffected)
	respondWithJSON(w, http.StatusOK, map[string]string{"result": response})
}

func (sc SeatsController) GetReservationSeats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	reservationId, err := strconv.Atoi(vars["reservationId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid reservation Id")
		return
	}

	seats, err := sc.SeatUsecases.GetUserSeats(reservationId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, seats)
}

func GetLayoutReservationData(vars map[string]string) (int, int, error) {
	layoutId, err := strconv.Atoi(vars["layoutId"])
	if err != nil {
		return 0, 0, err
	}

	sectionId, err := strconv.Atoi(vars["sectionId"])
	if err != nil {
		return 0, 0, err
	}

	return layoutId, sectionId, err
}

func (sc SeatsController) GetLayoutSeats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	layoutId, sectionId, err := GetLayoutReservationData(vars)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	seats, err := sc.SeatUsecases.GetLayoutSeats(layoutId, sectionId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, seats)
}
