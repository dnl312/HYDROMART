package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"merchant/service"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gomail "gopkg.in/mail.v2"
)

// SendMail sends top up link message
func SendMail(senderMail, destinationMail, ProductID string, quantity int, amount float64, mb service.MessageBroker) error {
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers with multiple recipients
	message.SetHeader("From", senderMail)
	message.SetHeader("To", destinationMail)
	message.SetHeader("Subject", "Transaction Processed")

	// Set email body
	message.SetBody("text/html", fmt.Sprintf(`
        <html>
            <body>
                <h1>Transaction Processed</h1>
                <p>The following order has been processed by our merchant.</p>
				<p><strong>Product ID: </strong> %s<br>
				<strong>Quantity: </strong>%d</p>
				<strong>Amount: </strong>$%.2f</a></p>
                <p>Thank you,<br>HydroMart</p>
            </body>
        </html>
		
    `, ProductID, quantity, amount,
	))

	emailDetails := struct {
		SenderMail      string  `json:"sender_mail"`
		DestinationMail string  `json:"destination_mail"`
		ProductID       string  `json:"product_id"`
		Quantity        int     `json:"quantity"`
		Amount          float64 `json:"amount"`
	}{
		SenderMail:      senderMail,
		DestinationMail: destinationMail,
		ProductID:       ProductID,
		Quantity:        quantity,
		Amount:          amount,
	}

	dataJson, err := json.Marshal(emailDetails)
	if err != nil {
		log.Printf("error marshalling email details: %v", err)
		return status.Error(codes.Internal, err.Error())
	}

	if err := mb.PublishMessage(dataJson); err != nil {
		return err
	}

	// Set up the SMTP dialer
	mailtrapPort, err := strconv.Atoi(os.Getenv("MAILTRAP_PORT"))
	if err != nil {
		return fmt.Errorf("sending top up email: %w", err)
	}
	dialer := gomail.NewDialer(
		os.Getenv("MAILTRAP_HOST"),
		mailtrapPort,
		os.Getenv("MAILTRAP_USERNAME"),
		os.Getenv("MAILTRAP_PASSWORD"),
	)
	// Send the email
	return dialer.DialAndSend(message)
}
