package utils

import (
	"fmt"
	"io"
	"net/http"
)

func FetchTemplateURL(url *string) (string, error) {
	resp, err := http.Get(*url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch template: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
