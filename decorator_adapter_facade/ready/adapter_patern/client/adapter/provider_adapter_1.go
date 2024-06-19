package adapter

import (
	"school/ready/adapter_patern/client"
	"school/ready/adapter_patern/domain"
	"time"
)

type (
	Provider1Adapter struct {
		sender *client.Provider1Sender
		// other specific services
	}
)

func NewProvider1Adapter(sender *client.Provider1Sender) *Provider1Adapter {
	return &Provider1Adapter{sender: sender}
}

func (e *Provider1Adapter) UpdateUser(user *domain.User) error {
	emailUser := client.P1User{
		Name:              user.Name,
		Age:               time.Now().Year() - user.BirthDate.Year(),
		Email:             user.Email,
		ProviderProjectId: e.GetProviderProjectId(user),
	}

	return e.sender.UpdateUser(&emailUser)
}

func (e *Provider1Adapter) GetProviderProjectId(user *domain.User) string {
	return "cool_" + user.ProjectId
}
