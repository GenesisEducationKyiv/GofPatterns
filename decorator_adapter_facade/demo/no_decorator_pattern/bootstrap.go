package decorator_pattern

import (
	"school/demo/no_decorator_pattern/client/finance"
	"school/demo/no_decorator_pattern/client/notification"
	"school/demo/no_decorator_pattern/domain"
	"school/demo/no_decorator_pattern/service"
)

func NewApplication() {
	rateFetcher := finance.NewRateFetcher()
	rateCache := finance.NewRateCache(rateFetcher)

	baseSender := notification.NewBaseSender()
	emailSender := notification.NewEmailSender(baseSender)
	smsSender := notification.NewSMSSender(emailSender)

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

	notificationService.SendNotification(
		&domain.User{
			Name: "John Doe",
		},
	)
}
