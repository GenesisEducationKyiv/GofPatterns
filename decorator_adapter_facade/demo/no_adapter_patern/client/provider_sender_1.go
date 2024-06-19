package client

type (
	Provider1Sender struct {
		// token
	}

	P1User struct {
		Name              string
		Age               int
		Email             string
		ProviderProjectId string
	}
)

func NewProvider1Sender() *Provider1Sender {
	return &Provider1Sender{}
}

func (s *Provider1Sender) UpdateUser(c *P1User) error {
	// send email process
	return nil
}
