package cli

import (
	"fmt"
	"strings"
)

func parseHeaders(headerList []string) (map[string]string, error) {
	headers := make(map[string]string)

	for _, h := range headerList {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header format: %s", h)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		headers[key] = value
	}

	return headers, nil
}