package notification

import (
	"school/ready/facade_pattern/domain"
	"school/ready/facade_pattern/service"
)

type (
	EmailSender struct {
		decorator service.NotificationSenderInterface
		// token
	}
)

func NewEmailSender(decorator service.NotificationSenderInterface) *EmailSender {
	return &EmailSender{
		decorator: decorator,
	}
}

func (s *EmailSender) Notify(user *domain.User) error {
	// send email process

	// Next step
	return s.decorator.Notify(user)
}
