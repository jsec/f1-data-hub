package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.helloWorldHandler)
	mux.HandleFunc("GET /seasons/{year}/standings", s.driverStandingsByYearHandler)

	return mux
}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"hello": "world",
	}

	s.respond(w, http.StatusOK, res)
}
