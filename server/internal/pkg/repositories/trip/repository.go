package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{
		db: db,
	}
}

func (r *TripRepository) GetTripAll() ([]entities.Trip, error) {
	var trips []entities.Trip

	rows, err := r.db.Query("SELECT * FROM pm.trip;")
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var trip entities.Trip
		if err := rows.Scan(&trip.ID, &trip.Name, &trip.Budget, &trip.CountryId, &trip.Description, &trip.StartDateTime, &trip.EndDateTime); err != nil {
			log.Printf("failed to scan row: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}
		trips = append(trips, trip)
	}

	if err := rows.Err(); err != nil {
		log.Printf("failed to scanning rows: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return trips, nil
}