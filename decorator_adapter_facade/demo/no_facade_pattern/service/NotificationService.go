package service

import (
	"school/demo/no_facade_pattern/domain"
)

type (
	NotificationService struct {
		sender     NotificationSenderInterface
		validation ValidationInterface
	}

	NotificationSenderInterface interface {
		Notify(user *domain.User) error
	}

	ValidationInterface interface {
		validateEmail(email string) bool
	}
)

func NewNotificationService(
	sender NotificationSenderInterface,
	validation ValidationInterface,
) *NotificationService {
	return &NotificationService{
		sender:     sender,
		validation: validation,
	}
}

func (s *NotificationService) SendNotification(user *domain.User) error {
	return s.sender.Notify(user)
}
