package main

import (
	"fmt"
	"net/http"
	"time"
)

// DoRequestWithRetry performs an HTTP request with retry and exponential backoff for 5xx errors.
func DoRequestWithRetry(reqFunc func() (*http.Response, error)) (*http.Response, error) {
	maxRetries := 3
	backoff := 100 * time.Millisecond
	for i := 0; i < maxRetries; i++ {
		resp, err := reqFunc()
		if err != nil {
			if i == maxRetries-1 {
				return nil, err
			}
			time.Sleep(backoff)
			backoff *= 2
			continue
		}
		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			if i == maxRetries-1 {
				return resp, nil // return last response
			}
			resp.Body.Close()
			time.Sleep(backoff)
			backoff *= 2
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("unreachable")
} 