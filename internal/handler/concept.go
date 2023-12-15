package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GET /api/v1/concepts
func (h *Handler) GetConcepts(c echo.Context) error {
	concepts, err := h.repo.GetConcepts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, concepts)
}

// GET /api/v1/concepts/:conceptID
func (h *Handler) GetConcept(c echo.Context) error {
	conceptId, err := uuid.Parse(c.Param("conceptID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	concept, err := h.repo.GetConcept(conceptId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, concept)
}

// DELETE /api/v1/concepts/:conceptID
func (h *Handler) DeleteConcept(c echo.Context) error {
	conceptId, err := uuid.Parse(c.Param("conceptID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = h.repo.DeleteConcept(conceptId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}

// GET /api/v1/users/:playerID/concepts
func (h *Handler) GetConceptsByPlayerId(c echo.Context) error {
	playerId, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	concepts, err := h.repo.GetConceptsByPlayerId(playerId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, concepts)
}

// POST /api/v1/reset/concepts
func (h *Handler) ResetConcepts(c echo.Context) error {
	err := h.repo.ResetConcepts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
