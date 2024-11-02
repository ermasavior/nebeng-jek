package http_client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type RestTransport struct {
	Url     string
	Method  string
	Header  http.Header
	Payload interface{}
}

func HttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	otelTr := otelhttp.NewTransport(tr)
	client := &http.Client{Transport: otelTr, Timeout: 10 * time.Second}
	return client
}

func HttpClientUseProxy(appEnv, proxyURL string) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	proxyUrl, _ := url.Parse(proxyURL)
	tr.Proxy = http.ProxyURL(proxyUrl)
	otelTr := otelhttp.NewTransport(tr)
	client := &http.Client{Transport: otelTr, Timeout: 10 * time.Second}
	return client
}

func Request(ctx context.Context, client *http.Client, req RestTransport) (res *http.Response, err error) {
	var httpReq *http.Request
	var body *bytes.Buffer
	if req.Payload != nil {
		b, err := json.Marshal(req.Payload)
		if err != nil {
			return res, err
		}
		body = bytes.NewBuffer(b)
	}

	if body == nil {
		httpReq, err = http.NewRequestWithContext(ctx, req.Method, req.Url, nil)
	} else {
		httpReq, err = http.NewRequestWithContext(ctx, req.Method, req.Url, body)
	}

	if err != nil {
		return res, err
	}

	if req.Header != nil {
		httpReq.Header = req.Header
	}

	res, err = client.Do(httpReq)
	if err != nil {
		return res, err
	}

	return res, nil
}
