package go_przelewy24

type GOP24 interface {
	RegisterTransaction(request RegistrationParams) (RegistrationResponse, error)
	VerifyTransaction(params Notification) error
}

type Gop24Impl struct {
	Config GOP24Config
	Client ClientWrapper
}

func NewGOP24(merchantID uint, posID uint, crcKey string, reportKey string, sanbox bool) GOP24 {
	var server Server
	if sanbox {
		server = SandboxServer
	} else {
		server = ProductionServer
	}

	return Gop24Impl{
		Config: GOP24Config{
			MerchantId: merchantID,
			PosId:      posID,
			CrcKey:     crcKey,
			Server:     server,
			ReportKey:  reportKey,
		},
		Client: NewClient(),
	}
}
