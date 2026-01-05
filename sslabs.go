package main

import (
	"encoding/json"
	"net/http"
)

func fetchSSLStatus(domain string) (SSLResponse, error) {
	url := "https://api.ssllabs.com/api/v3/analyze?host=" + domain + "&all=done"

	resp, err := http.Get(url)
	if err != nil {
		return SSLResponse{}, err
	}
	defer resp.Body.Close()

	var result SSLResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return SSLResponse{}, err
	}

	return result, nil
}
