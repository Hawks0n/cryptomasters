package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cryptomasters/internal/models"
)

func FetchPrices(url string) (models.CryptoPrice, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var prices models.CryptoPrice
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return prices, nil
}

