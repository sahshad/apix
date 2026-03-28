package client

import (
	"github.com/sahshad/apix/internal/config"
	"github.com/sahshad/apix/internal/types"
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
	"strings"
	"time"
)

type APIClient struct {
	BaseURL   string
	AuthToken string
	Client    *http.Client
}

func NewClient(cfg *config.Config) (*APIClient, error) {
	return &APIClient{
		BaseURL:   cfg.BaseURL,
		AuthToken: cfg.AuthToken,
		Client:    &http.Client{},
	}, nil
}

func (c *APIClient) Request(method, endpoint string, body io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest(method, c.BaseURL+"/"+endpoint, body)
	if err != nil {
		return nil, 0, err
	}

	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}

func (c *APIClient) Get(endpoint string) (*types.Response, error) {
	return c.request("GET", endpoint, nil)
}

func (c *APIClient) Post(endpoint string, data string) (*types.Response, error) {
	return c.request("POST", endpoint, bytes.NewBuffer([]byte(data)))
}

func (c *APIClient) Delete(endpoint string) (*types.Response, error) {
	return c.request("DELETE", endpoint, nil)
}

func (c *APIClient) Put(endpoint string, data string) (*types.Response, error) {
	return c.request("PUT", endpoint, bytes.NewBuffer([]byte(data)))
}

func (c *APIClient) Patch(endpoint string, data string) (*types.Response, error) {
	return c.request("PATCH", endpoint, bytes.NewBuffer([]byte(data)))
}

func (c *APIClient) SaveToFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}

// func (c *APIClient) request(method, endpoint string, body io.Reader) (*Response, error) {
// 	url := strings.TrimRight(c.BaseURL, "/") + "/" + strings.TrimLeft(endpoint, "/")

// 	req, err := http.NewRequest(method, url, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if c.AuthToken != "" {
// 		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
// 	}

// 	start := time.Now()

// 	resp, err := c.Client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	duration := time.Since(start)

// 	return &Response{
// 		Method:     method,
// 		StatusCode: resp.StatusCode,
// 		Body:       data,
// 		Headers:    resp.Header,
// 		DurationMs: duration.Milliseconds(),
// 		Size:       len(data),
// 	}, nil
// }

func (c *APIClient) request(method, endpoint string, body io.Reader) (*types.Response, error) {
	url := strings.TrimRight(c.BaseURL, "/") + "/" + strings.TrimLeft(endpoint, "/")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}

	// 🔹 Timing variables
	var timing types.Timing
	var dnsStart, connStart, tlsStart, reqStart time.Time

	trace := &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(httptrace.DNSDoneInfo) {
			timing.DNS = time.Since(dnsStart).Milliseconds()
		},

		ConnectStart: func(_, _ string) {
			connStart = time.Now()
		},
		ConnectDone: func(_, _ string, _ error) {
			timing.TCP = time.Since(connStart).Milliseconds()
		},

		TLSHandshakeStart: func() {
			tlsStart = time.Now()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			timing.TLS = time.Since(tlsStart).Milliseconds()
		},

		GotFirstResponseByte: func() {
			timing.TTFB = time.Since(reqStart).Milliseconds()
		},
	}

	reqStart = time.Now()

	// Attach trace
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	start := time.Now()

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	duration := time.Since(start)

	return &types.Response{
		Method:     method,
		StatusCode: resp.StatusCode,
		Body:       data,
		Headers:    resp.Header,
		DurationMs: duration.Milliseconds(),
		Size:       len(data),
		Timing:     &timing,
	}, nil
}