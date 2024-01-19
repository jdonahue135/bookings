package dbrepo

import (
	"errors"
	"time"

	"github.com/jdonahue135/bookings-app/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the DB
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the db
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	dateCutoff, _ := time.Parse("2006-01-02", "2050-01-01")
	if start.Before(dateCutoff) {
		return true, nil
	}
	return false, errors.New("some error")
}

// SearchAvailabilityForAllRooms returns a slice of available rooms (if any) for available date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	errorDate, _ := time.Parse("2006-01-02", "2060-01-01")
	if start == errorDate {
		return rooms, errors.New("some error")
	}
	dateCutoff, _ := time.Parse("2006-01-02", "2040-01-02")
	if start.Before(dateCutoff) {
		rooms = append(rooms, models.Room{})
	}
	return rooms, nil

}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("some error")
	}
	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "jack@nimble.com" {
		return 0, "", errors.New("some error")
	}
	return 1, "", nil
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) GetReservationById(id int) (models.Reservation, error) {
	var res models.Reservation

	return res, nil
}

func (m *testDBRepo) UpdateReservation(res models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction
	return restrictions, nil
}

func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

func (m *testDBRepo) DeleteBlockById(id int) error {
	return nil
}
