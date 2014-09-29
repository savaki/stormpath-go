package directories

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/httpctx"
	"github.com/savaki/stormpath-go/auth"
)

type Directories struct {
	client       httpctx.HttpClient
	ctx          context.Context
	directoryUrl string
}

func FromUrl(apiKey auth.ApiKey, directoryUrl string) *Directories {
	authFunc := auth.Authenticator(apiKey)
	return &Directories{
		client:       httpctx.WithAuthFunc(authFunc),
		ctx:          context.Background(),
		directoryUrl: directoryUrl,
	}
}
