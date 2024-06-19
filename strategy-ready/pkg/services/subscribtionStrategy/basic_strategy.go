package subscribtionStrategy

import "github.com/GenesisEducationKyiv/main-project-go/pkg/models"

type BasicStrategy struct {
}

func (b *BasicStrategy) Handle(subscriber models.Subscriber) error {
    // do some staff

    if subscriber.IsReadyForVip == false {
        // decline this subscription
    }

    return nil
}
