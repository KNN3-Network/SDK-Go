package knn3

import (
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 30,
}

func doRequest(method string, route, apiKey string, params map[string]interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/api/%s", baseURL, route)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("auth-key", apiKey)
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			switch v.(type) {
			case []string:
				q[k] = append(q[k], v.([]string)...)
			default:
				q.Add(k, fmt.Sprint(v))
			}
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
