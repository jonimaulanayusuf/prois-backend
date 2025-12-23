package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"prois-backend/internal/config"
	"time"
)

func getPurchaseWebhook() string {
	config.LoadEnv()
	return config.GetEnv("PURCHASE_WEBHOOK_URL", "")
}

func SendPurchaseWebhook(payload any) error {
	url := getPurchaseWebhook()
	return SendWebhook(url, payload)
}

func SendWebhook(url string, payload any) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook failed with status %d", resp.StatusCode)
	}

	return nil
}
