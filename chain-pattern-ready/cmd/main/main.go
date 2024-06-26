package main

import (
	"os"

	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/clients/chain"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/server"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/server/handlers"

	"github.com/GenesisEducationKyiv/main-project-go/pkg/clients"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/database"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/email"
	"github.com/GenesisEducationKyiv/main-project-go/pkg/services"
)

func main() {
	envValues, err := getEnvironmentValues()
	if err != nil {
		log.Fatalf("Error fetching environment values: %v", err)
	}

	port := envValues["PORT"]

	smtpHost := envValues["SMTP_HOST"]
	smtpPort := envValues["SMTP_PORT"]

	senderEmail := envValues["EMAIL"]
	senderPassword := envValues["EMAIL_PASSWORD"]

	fromCurrency := envValues["FROM_CURRENCY"]
	toCurrency := envValues["TO_CURRENCY"]

	dbFileFolder := envValues["DB_FILE_FOLDER"]
	dbFileName := envValues["DB_FILE_NAME"]

	binanceURL := envValues["BINANCE_API_URL"]
	coingeckoURL := envValues["COINGECKO_API_URL"]

	file, err := prepareFile(dbFileFolder, dbFileName)
	if err != nil {
		log.Fatalf("Error preparing file: %v", err)
	}

	db := database.NewFileDatabase(file)
	if db == nil {
		log.Fatalf("Error creating database")
	}

	repository := email.NewRepository(db)
	mailSender := email.NewSender(senderEmail,
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		smtp.PlainAuth("", senderEmail, senderPassword, smtpHost))

	binanceRateGetter := clients.NewLoggingClient("binance",
		clients.NewBinanceClient(binanceURL, &http.Client{}))
	coingeckoRateGetter := clients.NewLoggingClient("coingecko",
		clients.NewCoingeckoClient(coingeckoURL, &http.Client{}))

	binanceChain := chain.NewBaseChain(binanceRateGetter)
	coingeckoChain := chain.NewBaseChain(coingeckoRateGetter)
	binanceChain.SetNext(coingeckoChain)

	subscribeService := services.NewSubscribeService(repository)
	rateService := services.NewRateService(fromCurrency, toCurrency, binanceChain)
	emailService := services.NewEmailService(mailSender, &email.HTMLMessageBuilder{})

	api := handlers.NewAPIHandlers(emailService, rateService, subscribeService)
	web := handlers.NewWebHandlers(emailService, rateService, subscribeService)

	s := server.NewServer(api, web)

	log.Printf("Service listens on port: %s", port)

	err = s.Start(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func getEnvironmentValues() (map[string]string, error) {
	requiredEnvVars := []string{
		"PORT",
		"SMTP_HOST",
		"SMTP_PORT",
		"EMAIL",
		"EMAIL_PASSWORD",
		"FROM_CURRENCY",
		"TO_CURRENCY",
		"DB_FILE_FOLDER",
		"DB_FILE_NAME",
		"BINANCE_API_URL",
		"COINGECKO_API_URL",
	}

	envValues := make(map[string]string)

	for _, envVar := range requiredEnvVars {
		value, exists := os.LookupEnv(envVar)
		if !exists {
			return nil, fmt.Errorf("environment variable %s is not set", envVar)
		}
		envValues[envVar] = value
	}

	return envValues, nil
}

func prepareFile(dbFileFolder string, dbFileName string) (*os.File, error) {
	err := os.MkdirAll(dbFileFolder, os.ModeDir)
	if err != nil {
		return nil, fmt.Errorf("error creating directory: %w", err)
	}

	file, err := os.OpenFile(dbFileFolder+"/"+dbFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return file, nil
}
