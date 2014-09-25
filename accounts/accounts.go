package accounts

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/auth"
	"github.com/savaki/stormpath-go/internal/httputil"
)

type Accounts struct {
	client httputil.HttpClient
	ctx    context.Context
}

func New(authFunc auth.AuthFunc) *Accounts {
	httpClient := httputil.NewClient(authFunc)
	return &Accounts{
		client: httpClient,
		ctx:    context.Background(),
	}
}

func (a *Accounts) CreateAccount(account Account) error {
	return nil
}
