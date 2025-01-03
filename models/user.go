package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
	
)


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
	hashedPassword := utils.GetHashPassword(u.Password)
	result, err := stmp.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) VerifyCredentials() error {
	
	query := `
	SELECT password, id
	FROM users
	WHERE email = ?
	`

	var storedPass string
	err := db.DB.QueryRow(query, u.Email).Scan(&storedPass, &u.ID)

	if err != nil {
		return errors.New("user not found with the email")
	}

	matchPass := utils.ComparePassword(storedPass, u.Password)

	if !matchPass {
		return errors.New("user entered invalid password")
	}

	return nil
}
