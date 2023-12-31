package gop24

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RafalManka/go-przelewy24/internal"
	"net/http"
)

func (gop gop24Impl) VerifyTransaction(params Notification) error {
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
		MerchantID uint   `json:"merchantId"`
		PosID      uint   `json:"posId"`
		SessionID  string `json:"sessionId"`
		Amount     uint   `json:"amount"`
		Currency   string `json:"currency"`
		OrderID    uint   `json:"orderId"`
		Sign       string `json:"sign"`
	}{
		MerchantID: gop.Config.MerchantID,
		PosID:      gop.Config.PosID,
		SessionID:  params.SessionID,
		Amount:     params.Amount,
		Currency:   params.Currency,
		OrderID:    params.OrderID,
		Sign:       internal.HashData(signData),
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

	// TODO: verify that the response payload is correct
	fmt.Printf("Response Body: %s\n", string(responseBody))
	return nil
}
