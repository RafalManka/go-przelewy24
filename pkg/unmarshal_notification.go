package pkg

import (
	"encoding/json"
)

func UnmarshalNotification(body []byte) (Notification, error) {
	var target Notification
	err := json.Unmarshal(body, &target)
	return target, err
}
