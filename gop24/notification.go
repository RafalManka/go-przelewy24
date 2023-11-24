package gop24

import "encoding/json"

type Notification struct {
	MerchantID   uint   `json:"merchantId"`
	PosID        uint   `json:"posId"`
	SessionID    string `json:"sessionId"`
	Amount       uint   `json:"amount"`
	OriginAmount uint   `json:"originAmount"`
	Currency     string `json:"currency"`
	OrderID      uint   `json:"orderId"`
	MethodID     uint   `json:"methodId"`
	Statement    string `json:"statement"`
	Sign         string `json:"sign"`
}

func UnmarshalNotification(body []byte) (Notification, error) {
	var target Notification
	err := json.Unmarshal(body, &target)
	return target, err
}
