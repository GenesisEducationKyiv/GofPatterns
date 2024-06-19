package adapter_patern

import (
	"school/ready/adapter_patern/client"
	"school/ready/adapter_patern/client/adapter"
	"school/ready/adapter_patern/service"
)

func NewApplication() {
	p1 := client.NewProvider1Sender()
	p2 := client.NewProvider2Sender()

	p1Adapter := adapter.NewProvider1Adapter(p1)
	p2Adapter := adapter.NewProvider2Adapter(p2)

	// NotificationService
	_ = service.NewNotificationService(p1Adapter, p2Adapter)
}
