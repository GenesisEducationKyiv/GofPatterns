package service

// ProviderService Провайдерів багато, тож тут одним сервісом не обійдешся, але ми спростимо для зрозумілості 😺
type ProviderService struct {
}

func NewProviderService() *ProviderService {
	return &ProviderService{}
}

func (p *ProviderService) IsEmailExist(email string) bool {
	// Check if email is exist in email provider like gmail, yahoo, etc

	return true
}
