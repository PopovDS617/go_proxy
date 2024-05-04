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

func (s *HTTPServer) New() {
	s.getConfig()

	r := mux.NewRouter().SkipClean(true)
	r.HandleFunc("/proxy/{url:.*}", Proxy).Methods(http.MethodGet, http.MethodPost)

	server := &http.Server{
		Addr:    s.Addr,
		Handler: r,
	}

	s.Server = server
}

func (s *HTTPServer) getConfig() {
	port := os.Getenv("HTTP_PORT")

	if port == "" {
		port = ":9000"
	} else {
		port = ":" + port
	}

	s.Addr = port
}
