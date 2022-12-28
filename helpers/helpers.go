package helpers

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func SendEmail(email_address, token string) bool {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	toList := []string{email_address}

	host := "smtp.gmail.com"
	port := "587"

	body := []byte(token)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		// handling the errors
		fmt.Println(err)
		return false
	}
	return true
}

func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
