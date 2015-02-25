package wpsite

import (
	"code.google.com/p/go.net/context"
	"net/http"
)

// WPSite has useful methods for work with wordpress site
type WPSite interface {
	IsOnline(ctx context.Context) bool
	GetRedirect(ctx context.Context, url string) (string, error)
}

type HttpClient struct {
	Url    string
	client *http.Client
}

func NewHttpClient(url string) *HttpClient {
	return &HttpClient{
		Url:    url,
		client: http.DefaultClient,
	}
}

func (t *HttpClient) SetClient(client *http.Client) {
	t.client = client
}

// WPSite methods
func (t *HttpClient) IsOnline(ctx context.Context) bool {
	_, err := t.client.Get(t.Url)
	return err == nil
}

func (t *HttpClient) GetRedirect(ctx context.Context, url string) (string, error) {
	resp, err := t.client.Get(url)
	if err != nil {
		return url, err
	}
	u, err := resp.Location()
	if err != nil {
		if err != http.ErrNoLocation {
			err = nil
		}
		return url, err
	}
	return u.String(), nil
}
