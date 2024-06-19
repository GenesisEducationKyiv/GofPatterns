package notification

import (
	"school/ready/facade_pattern/domain"
	"school/ready/facade_pattern/service"
)

type (
	SMSSender struct {
		decorator service.NotificationSenderInterface
		// token
	}
)

func NewSMSSender(decorator service.NotificationSenderInterface) *SMSSender {
	return &SMSSender{
		decorator: decorator,
	}
}

func (s *SMSSender) Notify(user *domain.User) error {
	// send sms process

	// Next step
	return s.decorator.Notify(user)
}
