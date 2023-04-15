package knn3

import (
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 30,
}

func doRequest(method string, route, apiKey string) (*http.Response, error) {
	url := fmt.Sprintf("%s/api/%s", baseURL, route)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("auth-key", apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
