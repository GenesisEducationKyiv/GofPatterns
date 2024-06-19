package subscribtionStrategy

import "github.com/GenesisEducationKyiv/main-project-go/pkg/models"

type VipStrategy struct {
}

func (v *VipStrategy) Handle(subscriber models.Subscriber) error {
	// send promo discount

	// do something else

	return nil
}
