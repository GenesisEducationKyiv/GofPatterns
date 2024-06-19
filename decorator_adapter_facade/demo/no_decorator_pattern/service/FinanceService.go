package service

import "school/demo/no_decorator_pattern/domain"

type (
	FinanceService struct {
		fetcher RateFetcherInterface
	}

	RateFetcherInterface interface {
		Fetch(from, to domain.Currency) (float32, error)
		IsCurrencyAvailableForTrading(curr domain.Currency) (bool, error)
	}
)

func NewFinanceService(fetcher RateFetcherInterface) *FinanceService {
	return &FinanceService{fetcher: fetcher}
}

func (s *FinanceService) FetchRate(from, to domain.Currency) error {
	_, err := s.fetcher.Fetch(from, to)

	return err
}
