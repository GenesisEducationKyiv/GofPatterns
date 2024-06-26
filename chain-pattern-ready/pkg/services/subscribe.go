package services

import (
	"errors"
	"fmt"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/email"
)

type EmailDatabase interface {
	AddEmail(string) error
	EmailList() []string
}

type SubscribeService struct {
	database EmailDatabase
}

func NewSubscribeService(database EmailDatabase) *SubscribeService {
	return &SubscribeService{database: database}
}

var ErrAlreadySubscribed = fmt.Errorf("email is already subscribed")

func (s *SubscribeService) Subscribe(candidateEmail string) error {
	err := s.database.AddEmail(candidateEmail)
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
