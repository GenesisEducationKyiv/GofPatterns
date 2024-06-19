package service

type ProviderChecker struct {
}

func NewProviderChecker() *ProviderChecker {
	return &ProviderChecker{}
}

func (p *ProviderChecker) CheckProvider(email string) bool {
	// Check if email provider of receiver is valid

	return true
}
