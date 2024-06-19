package services

import (
	"errors"
	"fmt"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/models"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/email"
)

type EmailDatabase interface {
	Save(models.Subscriber) error
	EmailList() []string
}

type SubscribtionStrategy interface {
	Handle(subscriber models.Subscriber) error
}

type SubscribeService struct {
	database EmailDatabase
}

func NewSubscribeService(database EmailDatabase) *SubscribeService {
	return &SubscribeService{database: database}
}

var ErrAlreadySubscribed = fmt.Errorf("email is already subscribed")

func (s *SubscribeService) Subscribe(subscriber models.Subscriber, strategy SubscribtionStrategy) error {
	err := strategy.Handle(subscriber)
	if err != nil {
		return err
	}

	err = s.database.Save(subscriber)
	if err != nil {
		if errors.Is(err, email.ErrAlreadyExists) {
			return ErrAlreadySubscribed
		}
		return err
	}

	return nil
}

func (s *SubscribeService) EmailList() []string {
	return s.database.EmailList()
}
