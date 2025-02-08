package utils

import (
	"fmt"
	"time"
)

// TransformData takes raw alert data and transforms it into a format suitable for sending to Slack.
func TransformData(cpuUsage float64) string {
	timestamp := time.Now().Format(time.RFC1123)

	transformedMessage := fmt.Sprintf("Alert: CPU usage is at %.2f%% as of %s", cpuUsage, timestamp)
	return transformedMessage
}
