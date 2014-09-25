package stormpath

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/auth"
	"github.com/savaki/stormpath-go/internal/httputil"
	. "github.com/savaki/stormpath-go/types"
)

type Client struct {
	client httputil.HttpClient
	ctx    context.Context
	tenant *Tenant
}

func WithTenant(apiKey auth.ApiKey, tenant *Tenant) *Client {
	authFunc := auth.Authenticator(apiKey)
	return &Client{
		client: httputil.NewClient(authFunc),
		tenant: tenant,
		ctx:    context.Background(),
	}
}

func WithCurrentTenant(apiKey auth.ApiKey) (*Client, error) {
	return WithTenant(apiKey, nil).WithCurrentTenant()
}

func (c *Client) WithContext(ctx context.Context) *Client {
	return &Client{
		client: c.client,
		tenant: c.tenant,
		ctx:    ctx,
	}
}

func (c *Client) WithTenant(tenant *Tenant) *Client {
	return &Client{
		client: c.client,
		tenant: tenant,
		ctx:    c.ctx,
	}
}

func (c *Client) WithCurrentTenant() (*Client, error) {
	tenant, err := c.CurrentTenant()
	if err != nil {
		return nil, err
	}

	return c.WithTenant(tenant), nil
}

func (c *Client) CurrentTenant() (*Tenant, error) {
	tenant := &Tenant{}
	err := c.client.Get(c.ctx, BaseUrl+"/tenants/current", nil, tenant)
	return tenant, err
}

func (c *Client) Applications() ([]Application, error) {
	result := struct {
		Items []Application `json:"items"`
	}{}

	err := c.client.Get(c.ctx, c.tenant.Applications.Href(), nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

func (c *Client) Directories() ([]Map, error) {
	result := struct {
		Items []Map `json:"items"`
	}{}

	err := c.client.Get(c.ctx, c.tenant.Directories.Href(), nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
