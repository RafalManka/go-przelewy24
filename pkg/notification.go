package pkg

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
