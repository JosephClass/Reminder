package mail

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

// Information for server that we want to send
type EmailData struct {
	ServerId          string
	ServerName        string
	ServerDateExpired time.Time
	ServerDesc        string
	ServerIPAddress   string
}

// Email information ,include email data, sender data and email,and receiver email
type EmailRequest struct {
	EmailData     EmailData
	ReceiverEmail string

	emailHost      string
	emailPort      string
	emailSender    string
	passwordSender string
}

// Get the variable from env file
func (e *EmailRequest) init() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	e.emailHost = os.Getenv("EMAILHOST")
	e.emailPort = os.Getenv("EMAILPORT")
	e.emailSender = os.Getenv("EMAILACCOUNT")
	e.passwordSender = os.Getenv("EMAILPASSWORD")

}

// This function will parsing data to html template
func (e *EmailRequest) ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

//This function only send email,need to initiate
func (e *EmailRequest) SendEmail(template string) error {
	//Insert data into variable
	data := EmailData{
		ServerId:          e.EmailData.ServerId,
		ServerName:        e.EmailData.ServerName,
		ServerDateExpired: e.EmailData.ServerDateExpired,
		ServerDesc:        e.EmailData.ServerDesc,
		ServerIPAddress:   e.EmailData.ServerIPAddress,
	}

	//Parsing data into html
	result, err := e.ParseTemplate(template, data)
	if err != nil {
		log.Println("Email Error : Cannot Parsing Email")
	}

	//Initiate Gomail and Email Information
	mail := gomail.NewMessage()
	mail.SetHeader("From", "Invosa System Server Information")
	mail.SetHeader("To", e.ReceiverEmail)
	mail.SetHeader("Subject", "Server Payment Reminder")
	mail.SetBody("text/html", result)
	// mail.Attach(filename) // if user want to send file through email uncomment this

	//Convert port number to integer
	port, _ := strconv.Atoi(e.emailPort)

	//Sending Email
	dialer := gomail.NewDialer(e.emailHost,port,e.emailSender,e.passwordSender)

	errDial := dialer.DialAndSend(mail)
	if errDial != nil {
		log.Println("Email Error : Cannot connect to mail server")
	}

	return errDial
}


//This function w
func(e *EmailRequest) SendingEmailVerification()  {
	var err error
	
	//Location of html template file
	template := "template/index.html"

	err = e.SendEmail(template)

	if err != nil {
		fmt.Println("Email Error: " + err.Error() )
		return
	} 

	fmt.Println("Succesfully Sending Email to" + e.ReceiverEmail)
}