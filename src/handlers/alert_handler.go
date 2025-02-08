package handlers

import (
	"context"
	"encoding/json"
	"go-lambda-app/src/services"
	"go-lambda-app/src/utils"
	"log"
	"os"
)

type Alert struct {
	Detail struct {
		CPUUtilization float64 `json:"CPUUtilization"`
	} `json:"detail"`
}

func HandleAlert(ctx context.Context, event json.RawMessage) {

	log.Printf("Event: %v", string(event))

	webhookURL := os.Getenv("SLACK_WEBHOOK_URL_CPU")
	if webhookURL == "" {
		log.Printf("SLACK_WEBHOOK_URL environment variable is not set")
		return
	}

	var alert Alert
	err := json.Unmarshal(event, &alert)
	if err != nil {
		log.Printf("Error unmarshalling alert: %v", err)
		return
	}

	log.Printf("Unmarshal: %v", alert.Detail)

	transformedData := utils.TransformData(alert.Detail.CPUUtilization)

	log.Printf("transformedData: %v", transformedData)

	err = services.SendToSlack(webhookURL, transformedData)
	if err != nil {
		log.Printf("Error sending to Slack: %v", err)
	}
}
