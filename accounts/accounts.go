package accounts

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/auth"
	"github.com/savaki/stormpath-go/internal/httputil"
	"net/url"
)

type Accounts struct {
	client         httputil.HttpClient
	ctx            context.Context
	applicationUrl string
}

func FromUrl(apiKey auth.ApiKey, applicationUrl string) *Accounts {
	authFunc := auth.Authenticator(apiKey)
	return &Accounts{
		client:         httputil.NewClient(authFunc),
		ctx:            context.Background(),
		applicationUrl: applicationUrl,
	}
}

func (a *Accounts) WithContext(ctx context.Context) *Accounts {
	return &Accounts{
		client:         a.client,
		ctx:            ctx,
		applicationUrl: a.applicationUrl,
	}
}

func New(authFunc auth.AuthFunc) *Accounts {
	httpClient := httputil.NewClient(authFunc)
	return &Accounts{
		client: httpClient,
		ctx:    context.Background(),
	}
}

func (a *Accounts) CreateAccount(account Account) (created *Account, err error) {
	created = &Account{}
	err = a.client.Post(a.ctx, a.applicationUrl+"/accounts", account, created)
	return
}

func (a *Accounts) Search(email string) (accounts []Account, err error) {
	accounts = []Account{}
	params := &url.Values{}
	params.Add("email", email)
	err = a.client.Get(a.ctx, a.applicationUrl+"/accounts", params, nil)
	return
}
