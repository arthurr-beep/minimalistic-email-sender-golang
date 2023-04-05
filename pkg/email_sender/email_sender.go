package email_sender

import (
	"fmt"
	"net/smtp"
	"sync"
)

// SendEmail sends a single email using the provided Config and Message structs.
func SendEmail(config Config, msg Message) error {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.SMTPHost)

	messageBody := msg.buildMessage()

	recipients := append(msg.To, msg.Cc...)
	recipients = append(recipients, msg.Bcc...)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", config.SMTPHost, config.SMTPPort), auth, msg.From, recipients, []byte(messageBody))
	return err
}

// SendEmailConcurrent sends multiple email messages concurrently using the provided Config and a list of Message pointers.
func SendEmailConcurrent(config Config, messages []*Message) []error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(messages))

	for _, msg := range messages {
		wg.Add(1)
		go func(msg *Message) {
			defer wg.Done()
			err := SendEmail(config, *msg)
			errChan <- err
		}(msg)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
