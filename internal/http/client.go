package http

import (
	"net/http"
	"time"
)

var HTTPClient = &http.Client{
	Timeout: 5 * time.Second,
	Transport: &http.Transport{
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}
