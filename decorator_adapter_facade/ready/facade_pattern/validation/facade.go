package validation

import (
	"fmt"
	"school/ready/facade_pattern/validation/service"
)

type Facade struct {
	autocorrection   *service.AutoCorrection
	providerChecker  *service.ProviderChecker
	validation       *service.DNSValidation
	blacklistChecker *service.BlacklistChecker
	providerService  *service.ProviderService
}

func NewFacade(
	autocorrection *service.AutoCorrection,
	providerChecker *service.ProviderChecker,
	validation *service.DNSValidation,
	blacklistChecker *service.BlacklistChecker,
	providerService *service.ProviderService,
) *Facade {
	return &Facade{
		autocorrection:   autocorrection,
		providerChecker:  providerChecker,
		validation:       validation,
		blacklistChecker: blacklistChecker,
		providerService:  providerService,
	}
}

func (f *Facade) ValidateEmail(email string) bool {
	corEmail := f.autocorrection.CorrectEmail(email)
	if corEmail != email {
		fmt.Println("Email was corrected to: ", corEmail)
	}

	if ok := f.providerChecker.CheckProvider(corEmail); !ok {
		return false
	}

	if ok := f.validation.ValidateMXRecord(corEmail); !ok {
		return false
	}

	if blocked := f.blacklistChecker.IsInBlackList(corEmail); blocked {
		return false
	}

	if ok := f.providerService.IsEmailExist(corEmail); !ok {
		return false
	}

	return true
}
