package go_przelewy24

type GOP24Config struct {
	MerchantId uint
	PosId      uint
	CrcKey     string
	Server     Server
	ReportKey  string
}
