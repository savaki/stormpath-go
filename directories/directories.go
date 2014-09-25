package directories

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/stormpath-go/auth"
	"github.com/savaki/stormpath-go/internal/httputil"
)

type Directories struct {
	client       httputil.HttpClient
	ctx          context.Context
	directoryUrl string
}

func FromUrl(apiKey auth.ApiKey, directoryUrl string) *Directories {
	authFunc := auth.Authenticator(apiKey)
	return &Directories{
		client:       httputil.NewClient(authFunc),
		ctx:          context.Background(),
		directoryUrl: directoryUrl,
	}
}
