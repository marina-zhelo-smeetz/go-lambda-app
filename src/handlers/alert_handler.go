package handlers

import (
    "context"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/yourusername/go-lambda-app/src/services"
    "github.com/yourusername/go-lambda-app/src/utils"
    "log"
)

type Alert struct {
    Detail struct {
        CPUUtilization float64 `json:"CPUUtilization"`
    } `json:"detail"`
}

func HandleAlert(ctx context.Context, event events.CloudWatchEvent) {
    var alert Alert
    err := json.Unmarshal(event.Detail, &alert)
    if err != nil {
        log.Printf("Error unmarshalling alert: %v", err)
        return
    }

    transformedData := utils.TransformData(alert.Detail.CPUUtilization)
    err = services.SendToSlack(transformedData)
    if err != nil {
        log.Printf("Error sending to Slack: %v", err)
    }
}