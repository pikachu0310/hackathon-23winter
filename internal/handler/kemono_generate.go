package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/api"
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
	"net/http"
)

func (h *Handler) generateKemonoPromptAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), err.Error()).SetInternal(err)
	}

	prompt, err := api.GenerateKemonoPrompt(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Prompt = prompt

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoImageAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	image, err := api.GenerateKemonoImage(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Image = image

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoDescriptionAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoDescription, err := api.GenerateKemonoDescription(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Description = kemonoDescription

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoStatusAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoStatus, err := api.GenerateKemonoStatus(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	updateKemono := &domains.Kemono{
		ID:      kemono.ID,
		MaxHp:   kemonoStatus.MaxHP,
		Hp:      kemonoStatus.MaxHP,
		Attack:  kemonoStatus.Attack,
		Defense: kemonoStatus.Defence,
	}

	err = h.repo.UpdateKemono(c.Request().Context(), updateKemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoCharacterChipAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoCharacterChip, err := api.GenerateKemonoCharacterChip(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	updateKemono := &domains.Kemono{
		ID:    kemono.ID,
		Kind:  kemonoCharacterChip.Kind,
		Color: kemonoCharacterChip.Color,
	}

	err = h.repo.UpdateKemono(c.Request().Context(), updateKemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateKemonoNameAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoName, err := api.GenerateKemonoName(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	updateKemono := &domains.Kemono{
		ID:   kemono.ID,
		Name: kemonoName,
	}

	err = h.repo.UpdateKemono(c.Request().Context(), updateKemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateBreedKemonoPromptAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent1, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent1ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent2, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent2ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	prompt, err := api.BreedKemonoPrompt(kemonoParent1, kemonoParent2)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Prompt = prompt

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateBreedKemonoImageAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	image, err := api.GenerateKemonoImage(kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Image = image

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateBreedKemonoDescriptionAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent1, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent1ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent2, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent2ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoDescription, err := api.BreedKemonoDescription(kemonoParent1, kemonoParent2, kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Description = kemonoDescription

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateBreedKemonoStatusAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent1, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent1ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent2, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent2ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoStatus, err := api.BreedKemonoStatus(kemonoParent1, kemonoParent2, kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.MaxHp = kemonoStatus.MaxHP
	kemono.Hp = kemonoStatus.MaxHP
	kemono.Attack = kemonoStatus.Attack
	kemono.Defense = kemonoStatus.Defence

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return nil
}

func (h *Handler) generateBreedKemonoConceptsAndUpdateKemono(c echo.Context, kemonoID uuid.UUID) error {
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent1, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent1ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent2, err := h.repo.GetKemono(c.Request().Context(), *kemono.Parent2ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoConcepts, err := api.BreedKemonoConcepts(kemonoParent1, kemonoParent2, kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemono.Concepts = kemonoConcepts.Text()

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	return nil
}
