package pkg

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func (gop Gop24Impl) ApplyAuth(req *http.Request) {
	auth := fmt.Sprintf("%d:%s", gop.Config.MerchantId, gop.Config.ReportKey)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
}
