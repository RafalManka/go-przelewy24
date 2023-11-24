package go_przelewy24

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (gop Gop24Impl) VerifyTransaction(params Notification) error {
	// TODO: verify params
	// TODO: Verify config
	signData := fmt.Sprintf(
		`{"sessionId":"%s","orderId":%d,"amount":%d,"currency":"%s","crc":"%s"}`,
		params.SessionID,
		params.OrderID,
		params.Amount,
		params.Currency,
		gop.Config.CrcKey,
	)

	payload := struct {
		MerchantId uint   `json:"merchantId"`
		PosId      uint   `json:"posId"`
		SessionId  string `json:"sessionId"`
		Amount     uint   `json:"amount"`
		Currency   string `json:"currency"`
		OrderId    uint   `json:"orderId"`
		Sign       string `json:"sign"`
	}{
		MerchantId: gop.Config.MerchantId,
		PosId:      gop.Config.MerchantId,
		SessionId:  params.SessionID,
		Amount:     params.Amount,
		Currency:   params.Currency,
		OrderId:    params.OrderID,
		Sign:       HashData(signData),
	}
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", gop.Config.GetBaseUrl()+"/api/v1/transaction/verify", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	gop.ApplyAuth(req)

	responseBody, err := gop.Client.Call(req)
	if err != nil {
		return err
	}

	fmt.Printf("Response Body: %s\n", string(responseBody))
	return nil
}
