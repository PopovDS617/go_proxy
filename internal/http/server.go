package http

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	Server *http.Server
	Addr   string
}

func NewHTTPServer() *HTTPServer {
	addr := getConfig()
	r := mux.NewRouter().SkipClean(true)
	r.HandleFunc("/proxy/{url:.*}", proxy).Methods(http.MethodGet, http.MethodPost)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &HTTPServer{
		Server: server,
		Addr:   addr,
	}

}

func getConfig() string {
	port := os.Getenv("HTTP_PORT")

	if port == "" {
		port = ":9000"
	} else {
		port = ":" + port
	}

	return port
}
