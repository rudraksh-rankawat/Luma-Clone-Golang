package utils

import "golang.org/x/crypto/bcrypt"

func GetHashPassword(password string) string {
	bytePassword := []byte(password)
    hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
    if err != nil {
        return ""
    }
    return string(hash)
}

func ComparePassword(hashPassword, password string) bool {

    err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
    return err == nil
}

