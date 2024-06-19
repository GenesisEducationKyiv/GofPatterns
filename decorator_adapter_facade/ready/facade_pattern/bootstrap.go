package facade_pattern

import (
	"school/ready/facade_pattern/client/finance"
	"school/ready/facade_pattern/client/notification"
	"school/ready/facade_pattern/domain"
	"school/ready/facade_pattern/service"
	"school/ready/facade_pattern/validation"
	validationService "school/ready/facade_pattern/validation/service"
)

func NewApplication() {
	rateFetcher := finance.NewRateFetcher()
	rateCache := finance.NewRateCache(rateFetcher)

	baseSender := notification.NewBaseSender()
	emailSender := notification.NewEmailSender(baseSender)
	smsSender := notification.NewSMSSender(emailSender)

	autoCorrection := validationService.AutoCorrection{}
	providerChecker := validationService.ProviderChecker{}
	dnsValidation := validationService.DNSValidation{}
	blacklistChecker := validationService.BlacklistChecker{}
	providerService := validationService.ProviderService{}

	validationFacade := validation.NewFacade(
		&autoCorrection,
		&providerChecker,
		&dnsValidation,
		&blacklistChecker,
		&providerService,
	)

	// NotificationService
	notificationService := service.NewNotificationService(smsSender, validationFacade)

	// RateService
	rateService := service.NewFinanceService(rateCache)

	// Application logic // Not bootstrap logic
	rateService.FetchRate(
		domain.Currency{
			ISO3: "USD",
		},
		domain.Currency{
			ISO3: "EUR",
		},
	)

	// NotificationService
	notificationService.SendNotification(
		&domain.User{
			Name: "John Doe",
		},
	)
}
