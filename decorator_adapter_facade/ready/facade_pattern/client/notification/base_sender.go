package notification

import "school/ready/facade_pattern/domain"

type BaseSender struct{}

func NewBaseSender() *BaseSender {
	return &BaseSender{}
}

func (s *BaseSender) Notify(user *domain.User) error {
	return nil
}
