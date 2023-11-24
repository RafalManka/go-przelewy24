package internal

func (c GOP24Config) GetBaseUrl() string {
	if c.Server == ProductionServer {
		return "https://secure.przelewy24.pl"
	}
	return "https://sandbox.przelewy24.pl"
}
