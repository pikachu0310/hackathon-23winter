package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/repository"
	"github.com/pikachu0310/hackathon-23winter/src/images"
	"net/http"
	"strconv"
)

// スキーマ定義
type (
	/*
		Kemono struct {
			ID            uuid.UUID `db:"id"`
			Image         []byte    `db:"image"`
			Prompt        string    `db:"prompt"`
			Name          string    `db:"name"`
			Description   string    `db:"description"`
			CharacterChip int       `db:"character_chip"`
			IsPlayer      bool      `db:"is_player"`
			PlayerID      uuid.UUID `db:"player_id"`
			IsOwned       bool      `db:"is_owned"`
			OwnerID       uuid.UUID `db:"owner_id"`
			IsInField     bool      `db:"is_in_field"`
			IsBoss        bool      `db:"is_boss"`
			Field         int       `db:"field"`
			X             int       `db:"x"`
			Y             int       `db:"y"`
			HasParent     bool      `db:"has_parent"`
			Parent1ID     uuid.UUID `db:"parent1_id"`
			Parent2ID     uuid.UUID `db:"parent2_id"`
			HasChild      bool      `db:"has_child"`
			ChildID       uuid.UUID `db:"child_id"`
			Hp            int       `db:"hp"`
			Attack        int       `db:"attack"`
			Defense       int       `db:"defense"`
			CreatedAt     string    `db:"created_at"`
		}
	*/

	GetKemonosResponse []GetKemonoResponse

	GetKemonoResponse struct {
		ID            uuid.UUID `json:"id"`
		Image         []byte    `json:"image"`
		Prompt        string    `json:"prompt"`
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		CharacterChip int       `json:"character_chip"`
		IsPlayer      bool      `json:"is_player"`
		PlayerID      uuid.UUID `json:"player_id"`
		IsOwned       bool      `json:"is_owned"`
		OwnerID       uuid.UUID `json:"owner_id"`
		IsInField     bool      `json:"is_in_field"`
		IsBoss        bool      `json:"is_boss"`
		Field         int       `json:"field"`
		X             int       `json:"x"`
		Y             int       `json:"y"`
		HasParent     bool      `json:"has_parent"`
		Parent1ID     uuid.UUID `json:"parent1_id"`
		Parent2ID     uuid.UUID `json:"parent2_id"`
		HasChild      bool      `json:"has_child"`
		ChildID       uuid.UUID `json:"child_id"`
		Hp            int       `json:"hp"`
		Attack        int       `json:"attack"`
		Defense       int       `json:"defense"`
		CreatedAt     string    `json:"created_at"`
	}
)

func kemonoToGetKemonoResponse(kemono *repository.Kemono) GetKemonoResponse {
	return GetKemonoResponse{
		ID:            kemono.ID,
		Image:         kemono.Image,
		Prompt:        kemono.Prompt,
		Name:          kemono.Name,
		Description:   kemono.Description,
		CharacterChip: kemono.CharacterChip,
		IsPlayer:      kemono.IsPlayer,
		PlayerID:      kemono.PlayerID,
		IsOwned:       kemono.IsOwned,
		OwnerID:       kemono.OwnerID,
		IsInField:     kemono.IsInField,
		IsBoss:        kemono.IsBoss,
		Field:         kemono.Field,
		X:             kemono.X,
		Y:             kemono.Y,
		HasParent:     kemono.HasParent,
		Parent1ID:     kemono.Parent1ID,
		Parent2ID:     kemono.Parent2ID,
		HasChild:      kemono.HasChild,
		ChildID:       kemono.ChildID,
		Hp:            kemono.Hp,
		Attack:        kemono.Attack,
		Defense:       kemono.Defense,
		CreatedAt:     kemono.CreatedAt,
	}
}

// GET /api/v1/kemonos
func (h *Handler) GetKemonos(c echo.Context) error {
	kemonos, err := h.repo.GetKemonos(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := make(GetKemonosResponse, len(kemonos))
	for i, kemono := range kemonos {
		res[i] = kemonoToGetKemonoResponse(&kemono)
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/kemonos/:kemonoID
func (h *Handler) GetKemono(c echo.Context) error {
	kemonoID, err := uuid.Parse(c.Param("kemonoID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemonoID").SetInternal(err)
	}

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := kemonoToGetKemonoResponse(kemono)

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/kemonos/fields/:fieldID
func (h *Handler) GetKemonosByField(c echo.Context) error {
	fieldID, err := strconv.Atoi(c.Param("fieldID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid fieldID").SetInternal(err)
	}

	kemonos, err := h.repo.GetKemonosByField(c.Request().Context(), fieldID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := make(GetKemonosResponse, len(kemonos))
	for i, kemono := range kemonos {
		res[i] = kemonoToGetKemonoResponse(&kemono)
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/kemonos
func (h *Handler) CreateKemono(c echo.Context) error {
	id1, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")
	id2, _ := uuid.Parse("00000000-0000-0000-0000-000000000002")
	id3, _ := uuid.Parse("00000000-0000-0000-0000-000000000003")

	kemono := &repository.Kemono{
		ID:            id1,
		Image:         images.TestKemonoImage,
		Prompt:        "player",
		Name:          "player",
		Description:   "player",
		CharacterChip: 0,
		IsPlayer:      true,
		PlayerID:      id1,
		IsOwned:       false,
		OwnerID:       uuid.Nil,
		IsInField:     true,
		IsBoss:        false,
		Field:         1,
		X:             10,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		Hp:            100,
		Attack:        10,
		Defense:       10,
	}
	createdKemonoUUID, err := h.repo.CreateKemono(c.Request().Context(), kemono)
	if err != nil {
		return err
	}

	kemono = &repository.Kemono{
		ID:            id2,
		Image:         images.TestKemonoImage,
		Prompt:        "test1",
		Name:          "test1",
		Description:   "test1",
		CharacterChip: 1,
		IsPlayer:      false,
		PlayerID:      uuid.Nil,
		IsOwned:       false,
		OwnerID:       uuid.Nil,
		IsInField:     true,
		IsBoss:        false,
		Field:         1,
		X:             10,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		Hp:            20,
		Attack:        5,
		Defense:       2,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemono)
	if err != nil {
		return err
	}

	kemono = &repository.Kemono{
		ID:            id3,
		Image:         images.TestKemonoImage,
		Prompt:        "test2",
		Name:          "test2",
		Description:   "test2",
		CharacterChip: 2,
		IsPlayer:      false,
		PlayerID:      uuid.Nil,
		IsOwned:       false,
		OwnerID:       uuid.Nil,
		IsInField:     true,
		IsBoss:        false,
		Field:         1,
		X:             10,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		Hp:            30,
		Attack:        7,
		Defense:       4,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemono)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdKemonoUUID)
}
