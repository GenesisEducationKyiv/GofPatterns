package service

import "school/ready/adapter_patern/domain"

type (
	NotificationService struct {
		sender []EmailSenderInterface
	}

	EmailSenderInterface interface {
		UpdateUser(user *domain.User) error
	}
)

func NewNotificationService(sender ...EmailSenderInterface) *NotificationService {
	return &NotificationService{sender: sender}
}

func (s *NotificationService) SendNotification(user *domain.User) error {
	for _, sender := range s.sender {
		if err := sender.UpdateUser(user); err != nil {
			return err
		}
	}

	return nil
}
