package models

import (
	"database/sql"

	"github.com/Roni6291/event_booking/utils"
)

type User struct {
	Id       int64
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save(db *sql.DB) error {
	query := `INSERT INTO users(
		name, 
		email, 
		password
	  ) 
	  VALUES 
		(?, ?, ?)`
	cur, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer cur.Close()

	hashedPwd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := cur.Exec(
		u.Name,
		u.Email,
		hashedPwd,
	)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	// u.Id = id
	return err
}
