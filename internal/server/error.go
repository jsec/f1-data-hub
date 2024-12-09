package server

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error string `json:"string"`
}

func (s *Server) badRequest(w http.ResponseWriter, msg string) {
	s.httpError(w, http.StatusBadRequest, msg)
}

func (s *Server) httpError(w http.ResponseWriter, code int, msg string) {
	response := response{
		Error: msg,
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(res)
}
