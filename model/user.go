package model

import (
	"time"

	"com.go/event_booking/db"
)

type User struct {
	Id        int64
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	CreatedAt time.Time
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email, password, created_at)
	VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password, time.Now())

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.Id = id

	return err
}
