package bad_code

import (
	"fmt"
	"time"
)

var emailServer = "smtp.example.com"
var smsApiKey = "secret-key-123"
var slackWebhook = "https://hooks.slack.com/xxx"

type NotificationService struct {
	logs []string
}

func (n *NotificationService) Send(notifType string, recipient string, message string,
	priority string, retryCount int, attachmentPath string, templateId string) error {

	// Logging mixed in
	n.logs = append(n.logs, fmt.Sprintf("[%s] Sending %s to %s", time.Now(), notifType, recipient))

	// Validation scattered everywhere
	if recipient == "" {
		return fmt.Errorf("no recipient")
	}
	if message == "" && templateId == "" {
		return fmt.Errorf("no message")
	}

	// Template handling mixed with sending logic
	finalMessage := message
	if templateId == "welcome" {
		finalMessage = "Welcome to our platform! " + message
	} else if templateId == "alert" {
		finalMessage = "⚠️ ALERT: " + message
	} else if templateId == "promo" {
		finalMessage = "🎉 Special Offer: " + message
	}

	// Giant if-else chain - violates Open/Closed
	if notifType == "email" {
		fmt.Printf("Connecting to %s...\n", emailServer)
		fmt.Printf("Sending email to %s: %s\n", recipient, finalMessage)
		if attachmentPath != "" {
			fmt.Printf("Attaching file: %s\n", attachmentPath)
		}
		if priority == "high" {
			fmt.Println("Marking as HIGH priority")
		}
		// Retry logic duplicated
		for i := 0; i < retryCount; i++ {
			fmt.Println("Retrying email...")
		}

	} else if notifType == "sms" {
		if len(finalMessage) > 160 {
			finalMessage = finalMessage[:160]
		}
		fmt.Printf("Using API key: %s\n", smsApiKey)
		fmt.Printf("Sending SMS to %s: %s\n", recipient, finalMessage)
		// Retry logic duplicated again
		for i := 0; i < retryCount; i++ {
			fmt.Println("Retrying SMS...")
		}
	} else if notifType == "slack" {
		fmt.Printf("Posting to Slack webhook: %s\n", slackWebhook)
		fmt.Printf("Channel %s: %s\n", recipient, finalMessage)
		// Different retry logic - inconsistent!
		if retryCount > 0 {
			fmt.Println("Slack retry enabled")
		}

	} else if notifType == "push" {
		fmt.Printf("Sending push notification to device %s: %s\n", recipient, finalMessage)

	} else {
		return fmt.Errorf("unknown notification type: %s", notifType)
	}

	// Analytics mixed in
	fmt.Printf("Analytics: sent %s notification\n", notifType)

	// Rate limiting logic jammed in here too
	if priority != "high" {
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (n *NotificationService) GetUserPreferences(userId int) string {
	if userId == 1 {
		return "email"
	}
	return "sms"
}

func (n *NotificationService) FormatPhoneNumber(phone string) string {
	return "+1" + phone
}

func (n *NotificationService) ValidateEmail(email string) bool {
	return email != ""
}

func (n *NotificationService) GenerateReport() {
	fmt.Println("=== Notification Report ===")
	for _, log := range n.logs {
		fmt.Println(log)
	}
}

func main() {
	svc := &NotificationService{}

	svc.Send("email", "user@example.com", "Hello!", "high", 3, "/path/to/file.pdf", "welcome")
	svc.Send("sms", "5551234567", "Your code is 1234", "normal", 2, "", "")
	svc.Send("slack", "#general", "Deploy complete", "low", 1, "", "alert")

	svc.GenerateReport()
}
