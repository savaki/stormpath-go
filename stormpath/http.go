package stormpath

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Http interface {
	Get(path string, params *url.Values, v interface{}) error
	Post(path string, body interface{}, v interface{}) error
}

type HttpClient struct {
	ApiKey ApiKey
	Client *http.Client
}

func (h *HttpClient) Get(path string, params *url.Values, v interface{}) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	return h.Do(req, v)
}

func (h *HttpClient) Post(path string, data interface{}, v interface{}) error {
	var body io.Reader

	if data != nil {
		buffer, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(buffer)
	}

	req, _ := http.NewRequest("POST", BaseUrl+path, body)
	req.Header.Set("Content-Type", "application/json")
	return h.Do(req, v)
}

func (h *HttpClient) Do(req *http.Request, v interface{}) error {
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(h.ApiKey.Id, h.ApiKey.Secret)

	resp, err := h.Client.Transport.RoundTrip(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusFound {
		location := resp.Header.Get("Location")
		return h.Get(location, nil, v)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("returned status code => %d", resp.StatusCode))
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return nil
}
