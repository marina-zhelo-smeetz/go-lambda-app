package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func SendToSlack(webhookURL string, message string) error {
	slackMessage := SlackMessage{Text: message}
	jsonData, err := json.Marshal(slackMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to send message to Slack: %s", resp.Status)
	}

	return nil
}
