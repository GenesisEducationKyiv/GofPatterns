package client

type (
	Provider2Sender struct {
		// token
	}

	P2User struct {
		Email    string
		ClientId string
		Country  string
	}
)

func NewProvider2Sender() *Provider2Sender {
	return &Provider2Sender{}
}

func (s *Provider2Sender) SendEmail(c *P2User) error {
	// send email process
	return nil
}
