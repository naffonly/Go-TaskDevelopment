package util

import (
	"bytes"
	"github.com/dstotijn/go-notion"
	"io"
	"net/http"
	"time"
)

type httpTransport struct {
	w io.Writer
}

func (t *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	res.Body = io.NopCloser(io.TeeReader(res.Body, t.w))

	return res, nil
}

func InitNotionApi(key string) *notion.Client {

	buf := &bytes.Buffer{}

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &httpTransport{w: buf},
	}
	client := notion.NewClient(key, notion.WithHTTPClient(httpClient))

	return client
}
