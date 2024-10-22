package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) driverStandingsByYearHandler(c echo.Context) error {
	standings, err := s.db.GetDriverStandingsByYear(c.Request().Context(), 2020)
	if err != nil {
		// TODO: build out error handling
		log.Fatal("Things went bonk:", err)
	}

	return c.JSON(http.StatusOK, standings)
}
