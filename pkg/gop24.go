package pkg

import "github.com/RafalManka/go-przelewy24/internal"

type GOP24 interface {
	InitializeTransaction(params TransactionParams) (TransactionResults, error)
}

func NewGOP24() GOP24 {
	return internal.GOP24Impl{}
}
