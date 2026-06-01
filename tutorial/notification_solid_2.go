package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type EventType string

type Event struct {
	Type     EventType
	UserID   string
	UserName string
	Email    string
	Data     map[string]string
}

type Notification interface {
	Send(ctx context.Context, event Event, title string, body string) error
}

type FormatMessage interface {
	Format(event Event) (string, string)
}

type UserSignupEvent struct{}
type OrderPlacedEvent struct{}
type PaymentFailedEvent struct{}

func (u *UserSignupEvent) Format(event Event) (subject string, body string) {
	subject = "Welcome!"
	body = fmt.Sprintf("Hi %s, thanks for signing up!", event.UserName)
	return
}
func (o *OrderPlacedEvent) Format(event Event) (subject string, body string) {
	subject = "Order Confirmed"
	body = fmt.Sprintf("Hi %s, your order #%s has been placed.",
		event.UserName, event.Data["order_id"])
	return
}
func (p *PaymentFailedEvent) Format(event Event) (subject string, body string) {
	subject = "Payment Issue"
	body = fmt.Sprintf("Hi %s, your payment for order #%s failed. Please retry.",
		event.UserName, event.Data["order_id"])
	return
}

type EmailNotification struct {
	smtpHost string
}
type SMSNotification struct{}
type PushNotification struct {
	apiKey string
}

func (e *EmailNotification) Send(ctx context.Context, event Event, title string, body string) error {
	fmt.Printf("EMAIL to %s: [%s] %s\n", event.Email, title, body)
	return nil
}

func (s *SMSNotification) Send(ctx context.Context, event Event, title string, body string) error {
	short := body
	if len(short) > 160 {
		short = short[:157] + "..."
	}
	fmt.Printf("SMS to %s: %s\n", event.Data["phone"], short)
	return nil
}

func (p *PushNotification) Send(ctx context.Context, event Event, title string, body string) error {
	fmt.Printf("PUSH to %s: [%s] %s\n", event.UserID, title, strings.TrimSpace(body))
	return nil
}

type NotificationLogger interface {
	LogEntry(context context.Context, event Event, channel, status string) error
}
type PostgresLogger struct {
	db *sql.DB
}

func (pg *PostgresLogger) LogEntry(context context.Context, event Event, channel, status string) error {
	_, err := pg.db.ExecContext(context,
		"INSERT INTO notification_log (user_id, event_type, channel, status) VALUES ($1, $2, $3, $4)",
		event.UserID, event.Type, "all", "sent",
	)
	return err
}

type EventConfig struct {
	channel   []Notification
	formatter FormatMessage
}

type NotificationDispatcher struct {
	logger   NotificationLogger
	registry map[EventType]EventConfig
}

func NewNotificationDispatcher(logger NotificationLogger) *NotificationDispatcher {
	return &NotificationDispatcher{
		registry: make(map[EventType]EventConfig),
		logger:   logger}
}

func (d *NotificationDispatcher) Dispatch(ctx context.Context, event Event) error {
	// Format message based on event type
	config, ok := d.registry[event.Type]
	if !ok {
		return fmt.Errorf("unknow event type: %s", event.Type)
	}
	var subject, body string = config.formatter.Format(event)
	for _, sender := range config.channel {
		err := sender.Send(ctx, event, subject, body)
		if err != nil {
			return err
		}
	}
	err := d.logger.LogEntry(ctx, event, "all", "sent")
	return err

}
