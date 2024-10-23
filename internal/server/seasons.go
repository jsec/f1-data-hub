package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (s *Server) driverStandingsByYearHandler(c echo.Context) error {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid year")
	}

	standings, err := s.db.GetDriverStandingsByYear(c.Request().Context(), int32(year))
	if err != nil {
		// TODO: build out error handling
		log.Fatal("Things went bonk:", err)
	}

	return c.JSON(http.StatusOK, standings)
}
