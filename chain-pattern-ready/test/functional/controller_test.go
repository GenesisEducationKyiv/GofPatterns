package functional_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/models"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/database"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/email"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/server"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/server/handlers"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/services"
	"github.com/stretchr/testify/require"
)

func TestRate(t *testing.T) {
	t.Parallel()

	filename := "test-rate.db"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err = file.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(filename)

	s := prepareServer(t, file)

	testServer := httptest.NewServer(s)

	defer testServer.Close()

	serverURL := testServer.URL + "/api/rate"
	resp, err := http.Get(serverURL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	rate, err := strconv.ParseFloat(string(body[:len(body)-1]), 64)
	require.NoError(t, err)
	require.Equal(t, float64(10000), rate)
}

func TestSubscribe(t *testing.T) {
	t.Parallel()

	filename := "test-subscribe.db"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err = file.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(filename)

	s := prepareServer(t, file)

	testServer := httptest.NewServer(s)

	defer testServer.Close()

	serverURL := testServer.URL + "/api/subscribe"
	formData := url.Values{}
	formData.Set("email", "test@email.com")

	resp, err := http.Post(serverURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = http.Post(serverURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	require.Equal(t, http.StatusConflict, resp.StatusCode)
}

type mockRateService struct{}

func (m *mockRateService) Rate(_ context.Context, _, _ string) (*models.Rate, error) {
	return models.NewRate("BTC", "USDT", 10000), nil
}

type mockEmailSender struct{}

func (m *mockEmailSender) SendEmail(_ string, _ []byte) error {
	return nil
}

func prepareServer(t *testing.T, file *os.File) *server.Server {
	db := database.NewFileDatabase(file)
	if db == nil {
		t.Fatal("Error creating database")
	}

	repository := email.NewRepository(db)
	rateGetter := &mockRateService{}

	subscribeService := services.NewSubscribeService(repository)
	rateService := services.NewRateService("", "", rateGetter)
	emailService := services.NewEmailService(&mockEmailSender{}, &email.HTMLMessageBuilder{})

	api := handlers.NewAPIHandlers(emailService, rateService, subscribeService)

	s := server.NewServer(api, nil)
	return s
}
