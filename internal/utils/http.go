package utils

import "net/http"

func CopyRequestHeaders(destination http.Header, original http.Header) {
	for key, values := range original {
		for _, val := range values {
			destination.Add(key, val)
		}
	}
}
