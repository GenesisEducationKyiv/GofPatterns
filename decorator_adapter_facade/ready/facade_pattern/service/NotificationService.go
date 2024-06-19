package service

import (
	"school/ready/facade_pattern/domain"
	domainError "school/ready/facade_pattern/error"
)

type (
	NotificationService struct {
		sender           NotificationSenderInterface
		validationFacade ValidationInterface
	}

	NotificationSenderInterface interface {
		Notify(user *domain.User) error
	}

	ValidationInterface interface {
		ValidateEmail(email string) bool
	}
)

func NewNotificationService(sender NotificationSenderInterface, facade ValidationInterface) *NotificationService {
	return &NotificationService{
		sender:           sender,
		validationFacade: facade,
	}
}

func (s *NotificationService) SendNotification(user *domain.User) error {
	return s.sender.Notify(user)
}

func (s *NotificationService) Subscribe(user *domain.User) error {
	if ok := s.validationFacade.ValidateEmail(user.Email); !ok {
		return domainError.ErrInvalidEmail{}
	}

	// Subscribe logic

	return nil
}
