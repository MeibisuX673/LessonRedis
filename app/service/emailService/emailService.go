package emailService

//import (
//	"bytes"
//	"gopkg.in/gomail.v2"
//	"html/template"
//	"log"
//	"os"
//	"strconv"
//)
//
//type EmailService struct {
//}
//
//func New() *EmailService {
//	return &EmailService{}
//}
//
//func (es *EmailService) SendRegistration(toEmail, code string, ) {
//
//	from := os.Getenv("FROM_MAIL")
//	password := os.Getenv("PASSWORD")
//
//	smtpHost := os.Getenv("SMTP_HOST")
//	smtpPort := os.Getenv("SMTP_PORT")
//
//	port, err := strconv.Atoi(smtpPort)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	msg := gomail.NewMessage()
//	msg.SetHeader("From", from)
//	msg.SetHeader("To", toEmail)
//	msg.SetHeader("Subject", "Регистрация")
//
//	messageRegistration := "./emailMessages/registration.tmpl"
//
//	tmpl, err := template.New(messageRegistration).ParseFiles(messageRegistration)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = tmpl.Execute(os.Stdout, dogs)
//	if err != nil {
//		panic(err)
//	}
//	//body, err := os.ReadFile(os.Getenv("DIR_EMAIL_MESSAGES") + "/registration.html")
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//
//	htmlMessage := bytes.NewBuffer(body).String()
//
//	msg.SetBody("text/html", htmlMessage)
//
//	dialer := gomail.NewDialer(smtpHost, port, from, password)
//
//	if err := dialer.DialAndSend(msg); err != nil {
//		panic(err)
//	}
//
//}
