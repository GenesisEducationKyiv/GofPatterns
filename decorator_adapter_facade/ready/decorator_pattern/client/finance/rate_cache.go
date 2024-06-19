package finance

import (
	"school/ready/decorator_pattern/domain"
	"school/ready/decorator_pattern/service"
)

type (
	RateCache struct {
		decorator service.RateFetcherInterface

		unsafeCache map[string]float32 // які проблеми такого кешування? 😺
	}
)

func NewRateCache(decorator service.RateFetcherInterface) *RateCache {
	return &RateCache{
		decorator: decorator,
	}
}

func (s *RateCache) Fetch(from, to domain.Currency) (float32, error) {
	// Check cache
	if rate, ok := s.GetCache(from, to); ok {
		return rate, nil
	}

	// Fetch rate
	rate, err := s.decorator.Fetch(from, to)
	if err != nil {
		return 0, err
	}

	// Set cache
	s.SetCache(from, to, rate)

	return rate, nil
}

func (s *RateCache) IsCurrencyAvailableForTrading(curr domain.Currency) (bool, error) {
	return s.decorator.IsCurrencyAvailableForTrading(curr)
}

func (s *RateCache) GetCache(from, to domain.Currency) (float32, bool) {
	rate, ok := s.unsafeCache[s.generateKey(from, to)]

	return rate, ok
}

func (s *RateCache) SetCache(from, to domain.Currency, rate float32) {
	s.unsafeCache[s.generateKey(from, to)] = rate
}

func (s *RateCache) generateKey(from, to domain.Currency) string {
	return from.ISO3 + "_" + to.ISO3
}
