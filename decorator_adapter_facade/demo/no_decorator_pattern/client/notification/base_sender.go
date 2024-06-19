package notification

import "school/demo/no_decorator_pattern/domain"

type BaseSender struct{}

func NewBaseSender() *BaseSender {
    return &BaseSender{}
}

func (b BaseSender) Notify(user *domain.User) error {
    return nil
}

func (b BaseSender) IsValid(user *domain.User) error {
    return nil
}
