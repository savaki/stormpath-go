package httputil

import (
	"bytes"
	"code.google.com/p/go.net/context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/savaki/stormpath-go/auth"
	"io"
	"net/http"
	"net/url"
)

const (
	userAgent = "stormpath-go:0.1"
)

type HttpClient interface {
	Get(ctx context.Context, path string, params *url.Values, v interface{}) error
	Post(ctx context.Context, path string, body interface{}, v interface{}) error
}

func NewClient(authFunc auth.AuthFunc) HttpClient {
	return &client{
		authFunc: authFunc,
	}
}

type client struct {
	authFunc auth.AuthFunc
}

func (h *client) Get(ctx context.Context, path string, params *url.Values, v interface{}) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	return h.Do(ctx, req, v)
}

func (h *client) Post(ctx context.Context, path string, data interface{}, v interface{}) error {
	var body io.Reader

	if data != nil {
		buffer, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(buffer)
	}

	req, _ := http.NewRequest("POST", path, body)
	req.Header.Set("Content-Type", "application/json")
	return h.Do(ctx, req, v)
}

type response struct {
	resp *http.Response
	err  error
}

func (h *client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	if h.authFunc != nil {
		h.authFunc(req)
	}

	tr := &http.Transport{}
	ch := make(chan response, 1)
	defer close(ch)

	go func() {
		resp, err := tr.RoundTrip(req)
		ch <- response{resp: resp, err: err}
	}()

	var resp *http.Response
	var err error

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-ch
		return ctx.Err()
	case r := <-ch:
		resp = r.resp
		err = r.err
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusFound {
		location := resp.Header.Get("Location")
		return h.Get(ctx, location, nil, v)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("returned status code => %d", resp.StatusCode))
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return nil
}
