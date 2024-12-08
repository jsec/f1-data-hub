package server

import (
	"net/http"
	"strconv"
)

func (s *Server) driverStandingsByYearHandler(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(r.PathValue("year"))
	if err != nil {
		s.badRequest(w, "Invalid year")
	}

	standings, err := s.db.GetDriverStandingsByYear(r.Context(), int32(year))
	if err != nil {
		s.httpError(w, http.StatusInternalServerError, err.Error())
	}

	s.respond(w, http.StatusOK, standings)
}
