package smtp

import (
	"context"
	"database/sql"
	"errors"
	"exam/config"
	"exam/pkg"
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SendMail(toEmail string, msg string) error {

	// Compose the email message
	from := "amirjonqdirov28@gmail.com"
	to := []string{toEmail}
	subject := "Register for Customers"
	message := msg

	// Create the email message
	body := "To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + message

	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpServer)

	// Connectin to the SMTP server
	err := smtp.SendMail(config.SmtpServer+":"+config.SmtpPort, auth, from, to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}


func ExternalIdGenrerator(db *pgxpool.Pool) (string, error) {
	var (MaxId sql.NullString
	     maxId string)
	query := `SELECT MAX(external_id) FROM customers`
	err := db.QueryRow(context.Background(), query).Scan(&MaxId)
	maxId=pkg.NullNumberToString(MaxId)
	if err != nil {
		return "", err
	}
	nextId := 0
	currentMax, err := strconv.Atoi(maxId)
	if err != nil {
		fmt.Println(maxId)
		return maxId, errors.New("failed convert string to int")
	}
	nextId = currentMax + 1
	
	return fmt.Sprintf("%06d", nextId), nil
}
