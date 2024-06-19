package adapter

import (
	"school/ready/adapter_patern/client"
	"school/ready/adapter_patern/domain"
)

type (
	Provider2Adapter struct {
		sender *client.Provider2Sender
		// other specific services
	}
)

const ClientId = "It is really cool client id"

func NewProvider2Adapter(sender *client.Provider2Sender) *Provider2Adapter {
	return &Provider2Adapter{sender: sender}
}

func (e *Provider2Adapter) UpdateUser(user *domain.User) error {
	emailUser := client.P2User{
		Email:    user.Email,
		ClientId: ClientId,
		Country:  user.Country,
	}

	return e.sender.SendEmail(&emailUser)
}
