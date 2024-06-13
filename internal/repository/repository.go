package repository

import (
	"time"

	"github.com/saim61/udemy_bed_and_breakfast/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	GetRoomByID(id int) (models.Room, error)
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) 
}
