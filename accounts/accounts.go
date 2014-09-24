package accounts

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/stormpath"
	"net/http"
)

type Accounts struct {
	apiKey stormpath.ApiKey
	client stormpath.Http
	ctx    context.Context
}

func New(apiKey stormpath.ApiKey) *Accounts {
	client := &http.Client{
		Transport: &http.Transport{},
	}
	httpClient := &stormpath.HttpClient{ApiKey: apiKey, Client: client}
	return &Accounts{
		apiKey: apiKey,
		client: httpClient,
		ctx:    context.Background(),
	}
}

func (a *Accounts) CreateAccount(account Account) error {
	return nil
}
