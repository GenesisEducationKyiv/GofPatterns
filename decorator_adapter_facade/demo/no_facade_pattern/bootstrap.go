package facade_pattern

import (
	"school/demo/no_facade_pattern/client/finance"
	"school/demo/no_facade_pattern/client/notification"
	"school/demo/no_facade_pattern/domain"
	"school/demo/no_facade_pattern/service"
	validation2 "school/demo/no_facade_pattern/validation/service"
	"school/ready/facade_pattern/validation"
)

func NewApplication() {
	rateFetcher := finance.NewRateFetcher()
	rateCache := finance.NewRateCache(rateFetcher)

	baseSender := notification.NewBaseSender()
	emailSender := notification.NewEmailSender(baseSender)
	smsSender := notification.NewSMSSender(emailSender)

	autoCorrection := validation2.AutoCorrection{}
	providerChecker := validation2.ProviderChecker{}
	dnsValidation := validation2.DNSValidation{}
	blacklistChecker := validation2.BlacklistChecker{}
	providerService := validation2.ProviderService{}

	facade := validation.NewFacade(
		autoCorrection,
		providerChecker,
		dnsValidation,
		blacklistChecker,
		providerService,
	)

	// NotificationService
	notificationService := service.NewNotificationService(smsSender)

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
