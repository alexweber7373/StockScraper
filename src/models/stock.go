package models

import (
	"time"

	"example.com/stock-scraper/db"
)

type Stock struct {
	ID       int64
	Ticker   string    `binding:"required"`
	DateTime time.Time `binding:"required"`
	Value    float64   `binding:"required"`
}

func (s *Stock) Save() error {
	query := `
	INSERT INTO stocks(ticker, dateTime, value) VALUES (?, ?, ?)
	`

	// we are using prepare statement here becuase this may be used a lot in future code
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// at this point stmt has been prepared
	defer stmt.Close()

	// create current time
	dateTime := time.Now()
	s.DateTime = dateTime

	result, err := stmt.Exec(s.Ticker, s.DateTime, s.Value)

	if err != nil {
		return err
	}

	// we are still missing the ID, so we have to get that from the result and attach it to the object
	// we can do this because we are using a pointer and a copy of the value itself
	id, err := result.LastInsertId()

	s.ID = id

	return err

}
