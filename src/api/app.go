package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AugustasCepas/Theater-Seating-Algorithm/api/controllers"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/domain/repositories"
	infrastucture "github.com/AugustasCepas/Theater-Seating-Algorithm/infrastucture/repositories"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/usecases"
	"github.com/gorilla/mux"
)

type Services struct {
	SeatController *controllers.SeatsController
	SeatUsecases   *usecases.SeatUsecases
	SeatRepository repositories.SeatRepository
}

type App struct {
	Router   *mux.Router
	DB       *sql.DB
	Services Services
}

func (a *App) Initialize(DB *sql.DB) {
	var err error

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter().PathPrefix("/api").Subrouter()
	a.initializeServices()
	a.initializeRoutersV1()
}

func (a *App) initializeServices() {
	a.Services.SeatRepository = infrastucture.NewSeatRepository(a.DB)
	a.Services.SeatUsecases = usecases.NewSeatUsecases(a.Services.SeatRepository)
	a.Services.SeatController = controllers.NewSeatController(*a.Services.SeatUsecases)
}

func (a *App) initializeRoutersV1() {

	v1 := a.Router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/layout/{layoutId:[0-9]+}/section/{sectionId:[0-9]+}/rank/{rankId:[0-9]+}/reserve", a.Services.SeatController.ReserveSeats).Methods("POST")
	v1.HandleFunc("/seats/{reservationId:[0-9]+}", a.Services.SeatController.GetReservationSeats).Methods("GET")
	v1.HandleFunc("/layout/{layoutId:[0-9]+}/section/{sectionId:[0-9]+}/seats", a.Services.SeatController.GetLayoutSeats).Methods("GET")

}
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
