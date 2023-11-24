package gop24

import (
	"github.com/RafalManka/go-przelewy24/internal"
)

type Client interface {
	RegisterTransaction(request RegistrationParams) (RegistrationResponse, error)
	VerifyTransaction(params Notification) error
}

func NewClient(merchantID uint, posID uint, crcKey string, reportKey string, sandbox bool) Client {
	var server internal.Server
	if sandbox {
		server = internal.SandboxServer
	} else {
		server = internal.ProductionServer
	}

	return gop24Impl{
		Config: internal.GOP24Config{
			MerchantID: merchantID,
			PosID:      posID,
			CrcKey:     crcKey,
			Server:     server,
			ReportKey:  reportKey,
		},
		Client: internal.NewClient(),
	}
}
