package models

import "example.com/rest-api/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)`

	stmp, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmp.Close()
	result, err := stmp.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}
