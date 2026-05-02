package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptrace"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sahshad/apix/internal/config"
	"github.com/sahshad/apix/internal/types"
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
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *APIClient) Get(endpoint string) (*types.Response, error) {
	return c.send("GET", endpoint, nil, nil)
}

func (c *APIClient) Post(endpoint string, data string) (*types.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.send("POST", endpoint, bytes.NewBuffer([]byte(data)), headers)
}

func (c *APIClient) Delete(endpoint string) (*types.Response, error) {
	return c.send("DELETE", endpoint, nil, nil)
}

func (c *APIClient) Put(endpoint string, data string) (*types.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.send("PUT", endpoint, bytes.NewBuffer([]byte(data)), headers)
}

func (c *APIClient) Patch(endpoint string, data string) (*types.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.send("PATCH", endpoint, bytes.NewBuffer([]byte(data)), headers)
}

func (c *APIClient) SaveToFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}

func (c *APIClient) Multipart(method string, endpoint string, files []string) (*types.Response, error) {
	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	for _, item := range files {

		parts := strings.SplitN(item, "=", 2)

		if len(parts) != 2 {
			return nil, fmt.Errorf(
				"invalid form-file: %s",
				item,
			)
		}

		fieldName := parts[0]
		path := parts[1]

		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		defer file.Close()

		part, err := writer.CreateFormFile(
			fieldName,
			filepath.Base(path),
		)

		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}

	return c.send(method, endpoint, &body, headers)
}

func (c *APIClient) send(method, endpoint string, body io.Reader, headers map[string]string) (*types.Response, error) {
	url := strings.TrimRight(c.BaseURL, "/") + "/" + strings.TrimLeft(endpoint, "/")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Timing variables
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
