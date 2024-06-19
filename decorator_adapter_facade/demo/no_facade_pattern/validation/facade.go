package validation

import "school/ready/facade_pattern/validation/service"

type Facade struct {
	autocorrection *service.AutoCorrection
	bc             *service.BlacklistChecker
	dnsVal         *service.DNSValidation
	provChec       *service.ProviderChecker
	provServ       *service.ProviderService
}

func NewFacade(
	autocorrection *service.AutoCorrection,
	bc *service.BlacklistChecker,
	dnsVal *service.DNSValidation,
	provChec *service.ProviderChecker,
	provServ *service.ProviderService,
) *Facade {
	return &Facade{
		autocorrection: autocorrection,
		bc:             &service.BlacklistChecker{},
		dnsVal:         &service.DNSValidation{},
		provChec:       &service.ProviderChecker{},
		provServ:       &service.ProviderService{},
	}
}

func (f *Facade) validateEmail(email string) bool {
	//
	f.autocorrection.CorrectEmail()

	return false
}
