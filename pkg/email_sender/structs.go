package email_sender

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

// Config holds the SMTP configuration for sending emails.
// Contains the SMTP host, port, and user credentials.
type Config struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
}

// Attachment represents a file attachment with fields for Filename, ContentType, and Data.
type Attachment struct {
	Filename    string
	ContentType string
	Data        []byte
}

// Message represents an email message with fields for From, To, Cc, Bcc, Subject, Body, and Attachments.
type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        string
	Attachments []Attachment
}

// SetFrom sets the "From" field for the Message.
func (msg *Message) SetFrom(from string) {
	msg.From = from
}

// SetTo sets the "To" recipients for the Message.
func (msg *Message) SetTo(to []string) {
	msg.To = to
}

// SetCc sets the "Cc" recipients for the Message.
func (msg *Message) SetCc(cc []string) {
	msg.Cc = cc
}

// SetBcc sets the "Bcc" recipients for the Message.
func (msg *Message) SetBcc(bcc []string) {
	msg.Bcc = bcc
}

// SetSubject sets the "Subject" field for the Message.
func (msg *Message) SetSubject(subject string) {
	msg.Subject = subject
}

// SetBody sets the "Body" content for the Message.
func (msg *Message) SetBody(body string) {
	msg.Body = body
}

// AttachFile reads the file from the provided filepath, creates an Attachment, and adds it to the Message.
func (msg *Message) AttachFile(filepath, contentType string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	filename := filepath[strings.LastIndex(filepath, "/")+1:]
	attachment := Attachment{Filename: filename, ContentType: contentType, Data: data}
	msg.Attachments = append(msg.Attachments, attachment)
	return nil
}

// buildMessage constructs the MIME message for the email, including the email body and attachments.
func (msg *Message) buildMessage() string {
	header := make(map[string]string)
	header["From"] = msg.From
	header["To"] = strings.Join(msg.To, ",")
	header["Cc"] = strings.Join(msg.Cc, ",")
	header["Subject"] = msg.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "multipart/mixed; boundary=BOUNDARY"

	var message strings.Builder
	for key, value := range header {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	message.WriteString("\r\n--BOUNDARY\r\n")
	message.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	message.WriteString("\r\n")
	message.WriteString(msg.Body)

	for _, attachment := range msg.Attachments {
		message.WriteString("\r\n\r\n--BOUNDARY\r\n")
		message.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", attachment.ContentType, attachment.Filename))
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", attachment.Filename))
		message.WriteString("\r\n")
		message.WriteString(chunkBase64(attachment.Data))
	}
	message.WriteString("\r\n--BOUNDARY--")

	return message.String()
}

// chunkBase64 converts the byte slice data to a base64-encoded string and breaks it into 76-character lines.
func chunkBase64(data []byte) string {
	raw := base64.StdEncoding.EncodeToString(data)
	var chunked strings.Builder
	for i, l := 0, len(raw); i < l; i++ {
		if i > 0 && i%76 == 0 {
			chunked.WriteString("\r\n")
		}
		chunked.WriteByte(raw[i])
	}
	chunked.WriteString("\r\n")
	return chunked.String()
}
