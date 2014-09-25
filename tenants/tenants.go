package tenants

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/auth"
	"github.com/savaki/stormpath-go/internal/httputil"
)

type Tenants struct {
	client httputil.HttpClient
	ctx    context.Context
}

func New(authFunc auth.AuthFunc) *Tenants {
	httpClient := httputil.NewClient(authFunc)
	return &Tenants{
		client: httpClient,
		ctx:    context.Background(),
	}
}

func (t *Tenants) WithContext(ctx context.Context) *Tenants {
	return &Tenants{
		client: t.client,
		ctx:    ctx,
	}
}

func (t *Tenants) CurrentTenant() (tenant *Tenant, err error) {
	tenant = &Tenant{}
	err = t.client.Get(t.ctx, auth.BaseUrl+"/tenants/current", nil, tenant)
	return
}

func (t *Tenants) Applications() ([]Application, error) {
	tenant, err := t.CurrentTenant()
	if err != nil {
		return nil, err
	}

	result := struct {
		Items []Application `json:"items"`
	}{}

	err = t.client.Get(t.ctx, tenant.Applications.Href(), nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
