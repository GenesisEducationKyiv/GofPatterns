package finance

import "school/ready/facade_pattern/domain"

type (
	RateFetcher struct {
		// token
	}
)

func NewRateFetcher() *RateFetcher {
	return &RateFetcher{}
}

func (s *RateFetcher) Fetch(from, to domain.Currency) (float32, error) {
	// fetch rate process
	return 0, nil
}

func (s *RateFetcher) IsCurrencyAvailableForTrading(curr domain.Currency) (bool, error) {
	// check currency process
	return true, nil
}
