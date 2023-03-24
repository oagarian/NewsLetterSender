package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"modules/internal/db"
	"github.com/joho/godotenv"
    "os"
	_ "github.com/go-sql-driver/mysql"
	"net/smtp"
	"strings"
)

func DatabaseConnect() *db.Queries{
	dbconn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/NEWSLETTER")
	if err != nil {
		fmt.Println("Error opening database connection")
		log.Fatal(err)
	}
	
	database := db.New(dbconn);
	fmt.Println("Part 1")
	return database;
}


func sendEmail(emails []string) {
	var subject string = "Subject of the email"
	var body string = "Body of the email"

	error := godotenv.Load()

    if error != nil {
        fmt.Println("Error on loading .env")
        os.Exit(1)
    }

	smtpServer := os.Getenv("SMTP_SERVER")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpEmail := os.Getenv("SMTP_EMAIL")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpServer)

	message := "From: " + smtpEmail + "\n" +
        "To: " + strings.Join(emails, ",") + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpEmail, emails, []byte(message))
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
	
	fmt.Println("Email sucessfully sent")

}
func main() {
	database := DatabaseConnect();
	emails, err := database.GetEmails(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	sendEmail(emails)

}