package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (s *Service) HandleInsertJSON(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		s.logger.Error("failed to read request body", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to read request body",
		})
	}
	defer c.Request().Body.Close()

	if !json.Valid(body) {
		s.logger.Error("invalid JSON provided", nil)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON",
		})
	}

	repo := s.dataRepo
	if err := repo.InsertJSONPayload(body); err != nil {
		s.logger.Error("failed to insert JSON", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to insert JSON",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "JSON inserted successfully",
	})
}
