package types

import "net/http"

type Response struct {
	Method     string
	StatusCode int
	Body       []byte
	Headers    http.Header
	DurationMs int64
	Size       int
	Timing     *Timing
}

type ResponseParams struct {
	Method      string
	Endpoint    string
	Status      int
	ContentType string
	Body        string
	Duration    int64
	Size        string
	Headers     map[string][]string
	Timing      *Timing
}

type Timing struct {
	DNS  int64
	TCP  int64
	TLS  int64
	TTFB int64
}
