package event

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

type Event struct {
	ID         int `json:"id"`
	db         *sql.DB
	Name       string    `json:"name"`
	Host       string    `json:"host"`
	Location   string    `json:"location"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	DressCode  string    `json:"dress-code"`
	Theme      string    `json:"theme"`
	Price      float64   `json:"price"`
	SignupLink string    `json:"signup-link"`
}

func New(db *sql.DB) *Event {
	return &Event{db: db}
}

func (e *Event) Scan(r io.Reader) error {
	err := json.NewDecoder(r).Decode(e)
	return err
}

// Insert inserts event and all of its fields to database where event ID is created.
// It returns the created ID and an error.
func (e *Event) Insert() (int, error) {
	query := `
		INSERT INTO events (name, host, location, start, end, dc, theme, price, signup)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := e.db.Exec(query, e.Name, e.Host, e.Location, e.Start, e.End, e.DressCode, e.Theme, e.Price, e.SignupLink)
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert query: %v", err)
	}

	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	e.ID = int(eventID)
	return e.ID, nil
}

func (e *Event) GetByID(id int) error {
	query := `
		SELECT name, host, location, start, end, dc, theme, price, signup  FROM events WHERE id = ?
	`
	var start, end string
	err := e.db.QueryRow(query, id).Scan(&e.Name, &e.Host, &e.Location, &start, &end, &e.DressCode, &e.Theme, &e.Price, &e.SignupLink)
	if errors.Is(err, sql.ErrNoRows) {
		return &ErrEventNotFound{Message: "event not found", ID: id}
	} else if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	e.Start, err = time.Parse("2006-01-02 15:04:05", start)
	if err != nil {
		return err
	}
	e.End, err = time.Parse("2006-01-02 15:04:05", end)
	return err
}

type ErrEventNotFound struct {
	Message string
	ID      int
}

func (e *ErrEventNotFound) Error() string {
	return fmt.Sprintf("%s: %d", e.Message, e.ID)
}

func (e *ErrEventNotFound) Is(err error) bool {
	if err == nil {
		return false
	}
	var errEventNotFound *ErrEventNotFound
	ok := errors.As(err, &errEventNotFound)
	return ok
}
