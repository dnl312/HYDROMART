package utils

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	gomail "gopkg.in/mail.v2"
)

// SendMail sends top up link message
func SendMail(senderMail, destinationMail, ProductID string, quantity int32, amount float64) error {
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
				<strong>Quantity: </strong><a href="%d">payment-link</a></p>
				<strong>Amount: </strong><a href="%.2f">payment-link</a></p>
                <p>Thank you,<br>Book Rent</p>
            </body>
        </html>
		
    `, ProductID, quantity, amount,
	))

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
