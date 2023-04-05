# Minimalistic Email Sender

Minimalistic Email Sender is a Go package that provides a simple way to send emails with attachments using the SMTP protocol. It also supports concurrent email sending using goroutines.

## Features

- Send plain text emails
- Send emails with file attachments
- Send multiple emails concurrently

## Installation

To use the Email Sender package in your Go project, simply import it:

``` 
go get "github.com/minimalistic_email_sender/pkg/email_sender"

```

## USAGE
1. First, create a Config struct with your SMTP server details:

```
config := email_sender.Config{
	SMTPHost: "smtp.example.com",
	SMTPPort: "587",
	Username: "your_username",
	Password: "your_password",
}
```

2. Next, create a Message struct with your email details:

```
msg := &emailsender.Message{}
msg.SetFrom("you@example.com")
msg.SetTo([]string{"recipient1@example.com", "recipient2@example.com"})
msg.SetSubject("Test email")
msg.SetBody("This is a test email.")

```

3. Attach files to the email (optional):

```
err := msg.AttachFile("path/to/attachment.txt", "text/plain")

if err != nil {
	log.Fatal("Error attaching file:", err)
}

```

4. Send the email:

```
err = emailsender.SendEmail(config, *msg)
if err != nil {
	log.Println("Error sending email:", err)
} else {
	log.Println("Email sent successfully.")
}

```

5. To send multiple emails concurrently, create a list of Message pointers and call the SendEmailConcurrent function:

```
messages := []*emailsender.Message{}
// Add messages to the list

errors := emailsender.SendEmailConcurrent(config, messages)

if len(errors) > 0 {
	for _, err := range errors {
		log.Println("Error sending email:", err)
	}
} else {
	log.Println("All emails sent successfully.")
}

```

For more details see source code.
