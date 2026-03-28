package cli

import (
	"encoding/json"
	"fmt"
	"strings"
)

func FormatSize(bytes int) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}
	if bytes < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%.2f MB", float64(bytes)/(1024*1024))
}

func PrettyJSON(body []byte) (string, error) {
	var prettyJSON interface{}

	if err := json.Unmarshal(body, &prettyJSON); err != nil {
		return "", err
	}

	formatted, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return "", err
	}

	return string(formatted), nil
}

func formatBody(body, contentType string) string {
	if strings.Contains(contentType, "application/json") {
		formatted, err := PrettyJSON([]byte(body))
		if err != nil {
			return body
		}
		return formatted
	}
	return body
}