package service

import (
	"school/demo/no_decorator_pattern/domain"
)

type (
	NotificationService struct {
		sender NotificationSenderInterface
		// should be more
	}

	NotificationSenderInterface interface {
		Notify(user *domain.User) error
		IsValid(user *domain.User) error
	}
)

func NewNotificationService(sender NotificationSenderInterface) *NotificationService {
	return &NotificationService{
		sender: sender,
	}
}

func (s *NotificationService) CheckIntegrationsIsValid(user *domain.User) error {
	// Check email

	return s.sender.IsValid(user)
}

func (s *NotificationService) SendNotification(user *domain.User) error {
	// Send email

	return s.sender.Notify(user)
}

//func (s *NotificationService) SendNotification(user *domain.User) error {
//	// Send email
//	s.emailSender.UpdateUser(user)
//
//	// Send SMS
//	s.smsSender.UpdateUser(user)
//
//	return nil
//}
