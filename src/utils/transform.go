package utils

import (
	"fmt"
	"time"
)

// TransformData takes raw alert data and transforms it into a format suitable for sending to Slack.
func TransformData(alertData map[string]interface{}) string {
	cpuUsage := alertData["cpu_usage"].(float64)
	timestamp := time.Now().Format(time.RFC1123)

	transformedMessage := fmt.Sprintf("Alert: CPU usage is at %.2f%% as of %s", cpuUsage, timestamp)
	return transformedMessage
}