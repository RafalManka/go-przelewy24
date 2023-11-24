package gop24

import (
	"github.com/RafalManka/go-przelewy24/internal"
)

type GOP24 interface {
	RegisterTransaction(request RegistrationParams) (RegistrationResponse, error)
	VerifyTransaction(params Notification) error
}

func NewGOP24(merchantID uint, posID uint, crcKey string, reportKey string, sanbox bool) GOP24 {
	var server internal.Server
	if sanbox {
		server = internal.SandboxServer
	} else {
		server = internal.ProductionServer
	}

	return gop24Impl{
		Config: internal.GOP24Config{
			MerchantId: merchantID,
			PosId:      posID,
			CrcKey:     crcKey,
			Server:     server,
			ReportKey:  reportKey,
		},
		Client: internal.NewClient(),
	}
}
