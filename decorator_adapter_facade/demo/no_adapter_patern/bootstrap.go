package no_adapter_patern

import (
	notification "school/demo/no_adapter_patern/client"
	"school/demo/no_adapter_patern/service"
)

func NewApplication() {
	p1 := notification.NewProvider1Sender()
	p2 := notification.NewProvider2Sender()

	// NotificationService
	_ = service.NewNotificationService(*p1, *p2)
}
