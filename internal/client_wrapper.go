package internal

import (
	"io"
	"net/http"
)

type ClientWrapper interface {
	Call(req *http.Request) ([]byte, error)
}

type clientWrapperImpl struct {
	Client *http.Client
}

func NewClient() ClientWrapper {
	return &clientWrapperImpl{
		Client: &http.Client{},
	}
}

func (client clientWrapperImpl) Call(req *http.Request) ([]byte, error) {
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
