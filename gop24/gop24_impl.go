package gop24

import "github.com/RafalManka/go-przelewy24/internal"

type gop24Impl struct {
	Config internal.GOP24Config
	Client internal.ClientWrapper
}
