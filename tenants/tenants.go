package tenants

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/stormpath"
	"net/http"
)

type Tenants struct {
	apiKey stormpath.ApiKey
	client stormpath.Http
	ctx    context.Context
}

func New(apiKey stormpath.ApiKey) *Tenants {
	client := &http.Client{
		Transport: &http.Transport{},
	}
	httpClient := &stormpath.HttpClient{ApiKey: apiKey, Client: client}
	return &Tenants{
		apiKey: apiKey,
		client: httpClient,
		ctx:    context.Background(),
	}
}

func (t *Tenants) WithContext(ctx context.Context) *Tenants {
	return &Tenants{
		apiKey: t.apiKey,
		client: t.client,
		ctx:    ctx,
	}
}

func (t *Tenants) CurrentTenant() (tenant *Tenant, err error) {
	tenant = &Tenant{}
	err = t.client.Get(stormpath.BaseUrl+"/tenants/current", nil, tenant)
	return
}
