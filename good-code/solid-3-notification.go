package good_code

import "fmt"

type PriorityType string

var (
	HIGH PriorityType = "High"
	LOW  PriorityType = "Low"
)

type NotificationContext struct {
	Recipient      string
	Message        string
	Priority       PriorityType
	AttachmentPath string
}

type NotificationType string

var (
	EMAIL NotificationType = "Email"
	SMS   NotificationType = "SMS"
	SLACK NotificationType = "SLACK"
)

type TemplateType string

var (
	WELCOME TemplateType = "Welcome"
	ALERT   TemplateType = "Alert"
	PROMO   TemplateType = "Promo"
	GENERIC TemplateType = "GENERIC"
)

type MessageFormatter interface {
	Format(message string) string
}

type WelcomeTemplate struct{}
type AlertTemplate struct{}
type PromoTemplate struct{}
type GenericTemplate struct{}

func (w *GenericTemplate) Format(message string) string {
	return message
}
func (w *WelcomeTemplate) Format(message string) string {
	return "Welcome to our platform! " + message
}

func (a *AlertTemplate) Format(message string) string {
	return "⚠️ ALERT: " + message
}

func (p *PromoTemplate) Format(message string) string {
	return "🎉 Special Offer: " + message
}

type AttachmentSupporter interface {
	SupportsAttachment() bool
}

type NotificationSender interface {
	Send(NotificationContext) error
}

type EmailNotification struct {
	EmailServer string
}

func (e *EmailNotification) Send(context NotificationContext) error {

	if context.AttachmentPath != "" {
		fmt.Println("Attaching:", context.AttachmentPath)
	}
	fmt.Println("Sending email to:", context.Recipient)
	fmt.Println("Sending Notification Via Email channel")
	return nil
}

func (e *EmailNotification) SupportsAttachment() bool {
	return true
}

type NotificationServiceConfig struct {
	ApiKey  string
	Webhook string
	BaseURl string
}

type SMSNotification struct {
	SMSApiKey string
}

func (s *SMSNotification) Send(context NotificationContext) error {
	fmt.Println("Sending Notification Via Email channel")
	return nil
}

type SlackNotification struct {
	SlackWebhook string
}

func (s *SlackNotification) Send(context NotificationContext) error {
	fmt.Println("Sending Notification Via Slack channel")
	return nil
}

var TemplateRegistry = map[TemplateType]func() MessageFormatter{
	WELCOME: func() MessageFormatter {
		return &WelcomeTemplate{}
	},
	ALERT: func() MessageFormatter {
		return &AlertTemplate{}
	},
	PROMO: func() MessageFormatter {
		return &PromoTemplate{}
	},
	GENERIC: func() MessageFormatter {
		return &GenericTemplate{}
	},
}

func TemplateFactory(tType TemplateType) MessageFormatter {
	if creator, ok := TemplateRegistry[tType]; ok {
		return creator()
	}
	return &GenericTemplate{}
}

var NotificationRegistry = map[NotificationType]func(NotificationServiceConfig) NotificationSender{
	EMAIL: func(config NotificationServiceConfig) NotificationSender {
		return &EmailNotification{EmailServer: config.BaseURl}
	},
	SMS: func(config NotificationServiceConfig) NotificationSender {
		return &SMSNotification{SMSApiKey: config.ApiKey}
	},
	SLACK: func(config NotificationServiceConfig) NotificationSender {
		return &SlackNotification{SlackWebhook: config.Webhook}
	},
}

func NotificationFactory(nType NotificationType, config NotificationServiceConfig) (NotificationSender, error) {
	if creator, ok := NotificationRegistry[nType]; ok {
		return creator(config), nil
	}
	return nil, fmt.Errorf("unsupported notification type")
}

type NotificationService struct {
	notification NotificationSender
	formatter    MessageFormatter
}

type RetrySender struct {
	Wrapped    NotificationSender
	RetryCount int
}

type LoggingSender struct {
	logger NotificationSender
}

func (l *LoggingSender) Send(ctx NotificationContext) error {
	fmt.Println("Sending notification")
	err := l.logger.Send(ctx)
	return err
}

func (l *LoggingSender) SupportsAttachment() bool {
	if supporter, ok := l.logger.(AttachmentSupporter); ok {
		return supporter.SupportsAttachment()
	}
	return false
}

func (r *RetrySender) Send(ctx NotificationContext) error {
	var err error
	for i := 0; i <= r.RetryCount; i++ {
		err = r.Wrapped.Send(ctx)
		if err == nil {
			return nil
		}
	}
	return err
}

func (r *RetrySender) SupportsAttachment() bool {
	if supporter, ok := r.Wrapped.(AttachmentSupporter); ok {
		return supporter.SupportsAttachment()
	}
	return false
}

func NewRetrySender(sender NotificationSender, retryCount int) NotificationSender {
	return &RetrySender{
		Wrapped:    sender,
		RetryCount: retryCount,
	}
}

func NewLoggingSender(sender NotificationSender) NotificationSender {
	return &LoggingSender{
		logger: sender,
	}
}

func (n *NotificationService) Send(ctx NotificationContext) error {

	if ctx.AttachmentPath != "" {
		if _, ok := n.notification.(AttachmentSupporter); !ok {
			return fmt.Errorf("notification type dosen't support attachment")
		}
	}

	if n.formatter != nil {
		ctx.Message = n.formatter.Format(ctx.Message)
	}

	err := n.notification.Send(ctx)

	return err
}

func main() {

	messageFormatter := TemplateFactory(PROMO)
	config := NotificationServiceConfig{
		BaseURl: "https://chan.com",
	}
	notificationSender, err := NotificationFactory(EMAIL, config)
	wrapper := NewRetrySender(notificationSender, 0)
	loggingSender := NewLoggingSender(wrapper)
	if err != nil {
		panic(err)
	}

	notificationService := NotificationService{
		notification: loggingSender,
		formatter:    messageFormatter,
	}

	context := NotificationContext{
		Message:        "Hello There",
		Recipient:      "chandrakant@chandrakant.dev",
		Priority:       HIGH,
		AttachmentPath: "../../path.txt",
	}

	notificationService.Send(context)
}
