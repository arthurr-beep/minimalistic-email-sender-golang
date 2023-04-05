package email_sender

import (
	"strings"
	"testing"
)

func TestMessageBuild(t *testing.T) {
	msg := &Message{}
	msg.SetFrom("you@example.com")
	msg.SetTo([]string{"recipient1@example.com", "recipient2@example.com"})
	msg.SetSubject("Test email")
	msg.SetBody("This is a test email.")
	encodedMessage := msg.buildMessage()

	if !strings.Contains(encodedMessage, "From: you@example.com") {
		t.Error("BuildMessage() failed to set the From field")
	}

	if !strings.Contains(encodedMessage, "To: recipient1@example.com,recipient2@example.com") {
		t.Error("BuildMessage() failed to set the To field")
	}

	if !strings.Contains(encodedMessage, "Subject: Test email") {
		t.Error("BuildMessage() failed to set the Subject field")
	}

	if !strings.Contains(encodedMessage, "This is a test email.") {
		t.Error("BuildMessage() failed to set the Body field")
	}
}

func TestMessageAttachFile(t *testing.T) {
	msg := &Message{}
	err := msg.AttachFile("test_file.txt", "text/plain")
	if err != nil {
		t.Error("AttachFile() failed to attach a file")
	}

	if len(msg.Attachments) != 1 {
		t.Error("AttachFile() failed to add an attachment")
	}

	attachment := msg.Attachments[0]
	if attachment.Filename != "test_file.txt" || attachment.ContentType != "text/plain" {
		t.Error("AttachFile() failed to set attachment properties")
	}
}

func TestChunkBase64(t *testing.T) {
	data := []byte("This is a test string to be encoded in base64.")
	encoded := chunkBase64(data)

	if !strings.Contains(encoded, "VGhpcy") {
		t.Error("ChunkBase64() failed to encode data")
	}

	for _, line := range strings.Split(encoded, "\r\n") {
		if len(line) > 76 {
			t.Error("ChunkBase64() line exceeds 76 characters")
		}
	}
}
