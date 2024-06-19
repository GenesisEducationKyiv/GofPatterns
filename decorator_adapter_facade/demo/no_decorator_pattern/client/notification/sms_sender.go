package notification

import (
    "school/demo/no_decorator_pattern/domain"
    "school/demo/no_decorator_pattern/service"
)

type (
    SMSSender struct {
        decorator service.NotificationSenderInterface
        // token
    }
)

func NewSMSSender(decorator service.NotificationSenderInterface) *SMSSender {
    return &SMSSender{
        decorator: decorator,
    }
}

func (s *SMSSender) Notify(user *domain.User) error {
    // send sms process
    if !user.SMSSubscribed {
        return s.decorator.Notify(user)
    }

    // Next step
    return s.decorator.Notify(user)
}

func (s *SMSSender) IsValid(user *domain.User) error {
    // check phone is valid for now
    if !user.SMSSubscribed {
        return s.decorator.IsValid(user)
    }

    // Next step
    return s.decorator.IsValid(user)
}
