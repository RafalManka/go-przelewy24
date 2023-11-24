package pkg

import "github.com/RafalManka/go-przelewy24/internal"

type Gop24Impl struct {
	Config internal.GOP24Config
	Client internal.ClientWrapper
}
