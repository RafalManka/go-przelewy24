package go_przelewy24

import (
	"github.com/RafalManka/go-przelewy24/internal"
	"github.com/RafalManka/go-przelewy24/pkg"
)

type GOP24 interface {
	RegisterTransaction(request pkg.RegistrationParams) (pkg.RegistrationResponse, error)
	VerifyTransaction(params pkg.Notification) error
}

func NewGOP24(merchantID uint, posID uint, crcKey string, reportKey string, sanbox bool) GOP24 {
	var server internal.Server
	if sanbox {
		server = internal.SandboxServer
	} else {
		server = internal.ProductionServer
	}

	return pkg.Gop24Impl{
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
