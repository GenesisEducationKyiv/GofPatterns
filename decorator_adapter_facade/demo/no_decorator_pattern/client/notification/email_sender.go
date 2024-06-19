package notification

import (
    "school/demo/no_decorator_pattern/domain"
    "school/demo/no_decorator_pattern/service"
)

type (
    EmailSender struct {
        decorator service.NotificationSenderInterface
        // token
    }
)

func NewEmailSender(decorator service.NotificationSenderInterface) *EmailSender {
    return &EmailSender{
        decorator: decorator,
    }
}

func (s *EmailSender) Notify(user *domain.User) error {
    // send email process
    if !user.EmailSubscribed {
        return s.decorator.Notify(user)
    }

    // Next step
    return s.decorator.Notify(user)
}

func (s *EmailSender) IsValid(user *domain.User) error {
    // check email is valid for now
    if !user.EmailSubscribed {
        return s.decorator.IsValid(user)
    }

    // Next step
    return s.decorator.IsValid(user)
}
