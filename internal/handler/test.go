package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GET /api/v1/test
func (h *Handler) Test(c echo.Context) error {
	kemonoID, err := uuid.Parse("00000000-0000-0000-0000-000000000001")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemonoID").SetInternal(err)
	}

	fmt.Println(kemonoID)

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := kemonoToGetKemonoResponse(kemono)

	return c.JSON(http.StatusOK, res)
}
