package http

import (
	"errors"
	"fmt"

	"goproxy/internal/utils"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

func Proxy(w http.ResponseWriter, r *http.Request) {

	targetURL := mux.Vars(r)["url"]

	if targetURL == "" {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		slog.Info(fmt.Sprintf("%d :: %s", http.StatusBadRequest, "no url was provided"))
		return
	}

	if _, err := url.ParseRequestURI(targetURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Info(fmt.Sprintf("%d :: %s", http.StatusBadRequest, "url is not valid"))
		return
	}

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header = r.Header.Clone()

	resp, err := HTTPClient.Do(req)
	if err != nil {
		var nerr net.Error
		if errors.As(err, &nerr) {
			if nerr.Timeout() {
				http.Error(w, err.Error(), http.StatusGatewayTimeout)
				slog.Info(fmt.Sprintf("%d :: %s", http.StatusGatewayTimeout, targetURL))
				return
			}
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Info(fmt.Sprintf("%d :: %s", http.StatusInternalServerError, targetURL))
		return
	}

	defer resp.Body.Close()

	utils.CopyRequestHeaders(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)

	slog.Info(fmt.Sprintf("%d :: %s", resp.StatusCode, targetURL))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
