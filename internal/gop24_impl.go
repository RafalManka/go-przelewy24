package internal

import (
	"errors"
	"github.com/RafalManka/go-przelewy24/pkg"
)

type GOP24Impl struct {
}

func (G GOP24Impl) InitializeTransaction(params pkg.TransactionParams) (pkg.TransactionResults, error) {
	return pkg.TransactionResults{}, errors.New("Not implemented yet")
}
