package handlers

import (
	"context"
	"errors"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/services/subscribtionStrategy"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/models"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/services"
)

type (
	EmailService interface {
		SendEmails(context.Context, []string, *models.Message) error
	}

	RateService interface {
		Rate(ctx context.Context) (*models.Rate, error)
	}

	SubscribeService interface {
		Subscribe(subscriber models.Subscriber, strategy services.SubscribtionStrategy) error
		EmailList() []string
	}

	APIHandlers struct {
		emailService     EmailService
		rateService      RateService
		subscribeService SubscribeService
	}
)

func NewAPIHandlers(emailService EmailService,
	rateService RateService,
	subscribeService SubscribeService) *APIHandlers {
	return &APIHandlers{
		emailService:     emailService,
		rateService:      rateService,
		subscribeService: subscribeService,
	}
}

func (h *APIHandlers) Rate(c echo.Context) error {
	ctx := c.Request().Context()
	r, err := h.rateService.Rate(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, r.Rate())
}

func (h *APIHandlers) Subscribe(c echo.Context) error {
	email := strings.TrimSpace(c.FormValue("email"))

	subscriber := models.Subscriber{
		Email:            email,
		SubscriptionType: 1,
		IsReadyForVip:    false,
	}

	strategy := ChooseStrategy(subscriber)

	err := h.subscribeService.Subscribe(subscriber, strategy)
	if errors.Is(err, services.ErrAlreadySubscribed) {
		return c.JSON(http.StatusConflict, err.Error())
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	err = h.emailService.SendEmails(ctx, []string{email}, &models.Message{
		Subject: "Subscription",
		Body:    "You have successfully subscribed to the service",
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, email)
}

func (h *APIHandlers) SendEmails(c echo.Context) error {
	ctx := c.Request().Context()
	r, err := h.rateService.Rate(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.emailService.SendEmails(ctx, h.subscribeService.EmailList(), models.NewMessageFromRate(r))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Emails sent")
}

func ChooseStrategy(subscriber models.Subscriber) services.SubscribtionStrategy {
	switch subscriber.SubscriptionType {
	case 1:
		return &subscribtionStrategy.BasicStrategy{}
	case 2:
		return &subscribtionStrategy.VipStrategy{}
	}
}
