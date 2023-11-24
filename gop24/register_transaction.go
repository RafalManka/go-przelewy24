package gop24

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RafalManka/go-przelewy24/internal"
	"net/http"
)

type RegistrationParams struct {
	SessionId     string
	Amount        uint
	Currency      string
	Description   string
	Email         string
	FullName      string
	Country       string
	Language      string
	Phone         string
	AppointmentID uint
	UrlReturn     string
	UrlStatus     string
}

type RegistrationResponse struct {
	Token       string
	RedirectUrl string
}

func (gop gop24Impl) RegisterTransaction(request RegistrationParams) (RegistrationResponse, error) {
	// TODO: verify params
	// TODO: Verify config

	signData := fmt.Sprintf(
		`{"sessionId":"%s","merchantId":%d,"amount":%d,"currency":"%s","crc":"%s"}`,
		request.SessionId,
		gop.Config.MerchantId,
		request.Amount,
		request.Currency,
		gop.Config.CrcKey,
	)

	data := struct {
		MerchantID  uint   `json:"merchantId"`
		PosID       uint   `json:"posId"`
		SessionID   string `json:"sessionId"`
		Amount      uint   `json:"amount"`
		Currency    string `json:"currency"`
		Description string `json:"description"`
		Email       string `json:"email"`
		Client      string `json:"client"`
		Country     string `json:"country"`
		Phone       string `json:"phone"`
		Language    string `json:"language"`
		UrlStatus   string `json:"urlStatus"`
		UrlReturn   string `json:"urlReturn"`
		Sign        string `json:"sign"`
	}{
		MerchantID:  gop.Config.MerchantId,
		PosID:       gop.Config.PosId,
		SessionID:   request.SessionId,
		Amount:      request.Amount,
		Currency:    request.Currency,
		Description: request.Description,
		Email:       request.Email,
		Client:      request.FullName,
		Country:     request.Country,
		Phone:       request.Phone,
		Language:    request.Language,
		UrlReturn:   request.UrlReturn,
		UrlStatus:   request.UrlStatus,
		Sign:        internal.HashData(signData),
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return RegistrationResponse{}, err
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", gop.Config.GetBaseUrl()+"/api/v1/transaction/register", body)
	if err != nil {
		return RegistrationResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	gop.ApplyAuth(req)
	responseBody, err := gop.Client.Call(req)
	if err != nil {
		return RegistrationResponse{}, err
	}

	type tokenData struct {
		Token string `json:"token"`
	}
	type registrationResponse struct {
		Data         tokenData `json:"data"`
		ResponseCode uint      `json:"responseCode"`
	}
	var response registrationResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return RegistrationResponse{}, err
	}

	if response.Data.Token == "" {
		return RegistrationResponse{}, errors.New("error generating token")
	}

	tr := RegistrationResponse{
		Token:       response.Data.Token,
		RedirectUrl: fmt.Sprintf("%s/trnRequest/%s", gop.Config.GetBaseUrl(), response.Data.Token),
	}
	return tr, nil
}
