package gop24

import (
	"encoding/base64"
	"fmt"
	"github.com/RafalManka/go-przelewy24/internal"
	"net/http"
)

type gop24Impl struct {
	Config internal.GOP24Config
	Client internal.ClientWrapper
}

func (gop gop24Impl) ApplyAuth(req *http.Request) {
	auth := fmt.Sprintf("%d:%s", gop.Config.MerchantID, gop.Config.ReportKey)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
}
