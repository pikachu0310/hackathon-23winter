package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/api"
	"net/http"
)

func (h *Handler) generateKemonoPromptAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	prompt, err := api.GenerateKemonoPrompt(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	kemono.Prompt = prompt

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoImageAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	image, err := api.GenerateKemonoImage(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	kemono.Image = image

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoDescriptionAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	kemonoDescription, err := api.GenerateKemonoDescription(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	kemono.Description = kemonoDescription

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoStatusAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	kemonoStatus, err := api.GenerateKemonoStatus(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	kemono.MaxHp = kemonoStatus.MaxHP
	kemono.Hp = kemonoStatus.MaxHP
	kemono.Attack = kemonoStatus.Attack
	kemono.Defense = kemonoStatus.Defence

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoCharacterChipAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	kemonoCharacterChip, err := api.GenerateKemonoCharacterChip(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	kemono.Kind = kemonoCharacterChip.Kind
	kemono.Color = kemonoCharacterChip.Color

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoNameAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	kemonoName, err := api.GenerateKemonoName(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	kemono.Name = kemonoName

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return nil
}
