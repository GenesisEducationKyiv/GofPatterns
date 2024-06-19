package service

import "school/ready/decorator_pattern/domain"

type (
	NotificationService struct {
		sender NotificationSenderInterface
	}

	NotificationSenderInterface interface {
		Notify(user *domain.User) error
	}
)

func NewNotificationService(sender NotificationSenderInterface) *NotificationService {
	return &NotificationService{sender: sender}
}

func (s *NotificationService) SendNotification(user *domain.User) error {
	return s.sender.Notify(user)
}
