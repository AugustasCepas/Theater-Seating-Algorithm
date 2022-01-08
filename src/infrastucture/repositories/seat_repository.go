package infrastucture

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/AugustasCepas/Theater-Seating-Algorithm/domain/repositories"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/dtos"
)

type SeatRepositorySQL struct {
	DB *sql.DB
}

func NewSeatRepository(db *sql.DB) repositories.SeatRepository {
	return &SeatRepositorySQL{DB: db}
}

func (seatRepository SeatRepositorySQL) ReserveSeat(reservationId int, seatId int) (int64, error) {
	query := "UPDATE seats SET reservation_id=$1 WHERE id=$2"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := seatRepository.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, reservationId, seatId)
	if err != nil {
		log.Printf("Error %s when reservationg seat", err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected != 0 {
		log.Printf("(id: %d) seat reserved", seatId)
	}

	return rowsAffected, err
}

func (seatRepository SeatRepositorySQL) GetReservationSeats(reservationId int) ([]dtos.ReservedSeats, error) {
	var seats []dtos.ReservedSeats

	query := "SELECT seat_number, row_id FROM seats WHERE reservation_id=$1"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := seatRepository.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return seats, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, reservationId)
	if err != nil {
		log.Printf("Error %s when getting seats", err)
		return seats, err
	}

	seatsCount := 0
	for rows.Next() {
		var a dtos.ReservedSeats
		err := rows.Scan(&a.Number, &a.RowId)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		seats = append(seats, a)
		seatsCount += 1
	}

	if seatsCount > 0 {
		log.Printf("reservation_id: %d - %d seats received", reservationId, seatsCount)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return seats, nil
}

func (seatRepository SeatRepositorySQL) GetSectionSeats(layoutId int, sectionId int, rankId int) ([]dtos.SectionSeats, error) {
	var seats []dtos.SectionSeats

	query := `select se.id, se.rank_id, se.row_id, se.seat_index, se.seat_number from layouts l 	
	join sections sec on sec.layout_id = l.id
	join rows r on r.section_id = sec.id
	join seats se on se.row_id = r.id
	where l.id = $1 and sec.id = $2 and se.rank_id = $3 and se.reservation_id = 0
	order by r.id, se.seat_index`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := seatRepository.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return seats, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, layoutId, sectionId, rankId)
	if err != nil {
		log.Printf("Error %s when getting seats", err)
		return seats, err
	}

	seatsCount := 0
	for rows.Next() {
		var seat dtos.SectionSeats
		err := rows.Scan(&seat.Id, &seat.RankId, &seat.RowId, &seat.Index, &seat.Number)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		seats = append(seats, seat)
		seatsCount += 1
	}

	if seatsCount > 0 {
		log.Printf("sectionId: %d - %d seats received", sectionId, seatsCount)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return seats, nil
}

func (seatRepository SeatRepositorySQL) GetLastReservationId() (int, error) {
	var lastReservationId int

	err := seatRepository.DB.QueryRow("select max(reservation_id) from seats").Scan(&lastReservationId)

	if err != nil {
		return 0, err
	}

	log.Printf("Last reservation id: %d", lastReservationId)

	return lastReservationId, nil
}

func (seatRepository SeatRepositorySQL) GetLayoutSeats(layoutId int, sectionId int) ([]dtos.LayoutSeats, error) {
	var seats []dtos.LayoutSeats

	query := `select se.reservation_id from layouts l
	join sections sec on sec.layout_id = l.id 
	join rows r on r.section_id = sec.id 
	join seats se on se.row_id = r.id 
	where sec.id = $1 and l.id = $2
	order by r.id, se.seat_index `

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := seatRepository.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return seats, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, layoutId, sectionId)
	if err != nil {
		log.Printf("Error %s when getting seats", err)
		return seats, err
	}

	seatsCount := 0
	for rows.Next() {
		var seat dtos.LayoutSeats
		err := rows.Scan(&seat.ReservationId)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		seats = append(seats, seat)
		seatsCount += 1
	}

	if seatsCount > 0 {
		log.Printf("sectionId: %d - %d seats received", sectionId, seatsCount)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return seats, nil
}

func (seatRepository SeatRepositorySQL) GetLayoutRows(layoutId int, sectionId int) (int, error) {
	var layoutRows int

	query := `select count(layoutRows) from(
		select se.row_id as layoutRows from layouts l
		join sections sec on sec.layout_id = l.id
		join rows r on r.section_id = sec.id
		join seats se on se.row_id = r.id
		where sec.id = $1 and l.id = $2
		group by se.row_id) sub`

	err := seatRepository.DB.QueryRow(query, layoutId, sectionId).Scan(&layoutRows)

	if err != nil {
		return 0, err
	}

	log.Printf("Layout rows: %d", layoutRows)

	return layoutRows, nil
}
