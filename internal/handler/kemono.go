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
			ID            *uuid.UUID `db:"id"`
			Image         []byte     `db:"image"`
			Prompt        *string    `db:"prompt"`
			Concepts      *string    `db:"concepts"`
			Name          *string    `db:"name"`
			Description   *string    `db:"description"`
			CharacterChip *int       `db:"character_chip"`
			IsPlayer      *bool      `db:"is_player"`
			PlayerID      *uuid.UUID `db:"player_id"`
			IsOwned       *bool      `db:"is_owned"`
			OwnerID       *uuid.UUID `db:"owner_id"`
			IsInField     *bool      `db:"is_in_field"`
			IsBoss        *bool      `db:"is_boss"`
			Field         *int       `db:"field"`
			X             *int       `db:"x"`
			Y             *int       `db:"y"`
			HasParent     *bool      `db:"has_parent"`
			Parent1ID     *uuid.UUID `db:"parent1_id"`
			Parent2ID     *uuid.UUID `db:"parent2_id"`
			HasChild      *bool      `db:"has_child"`
			ChildID       *uuid.UUID `db:"child_id"`
			MaxHp         *int       `db:"max_hp"`
			Hp            *int       `db:"hp"`
			Attack        *int       `db:"attack"`
			Defense       *int       `db:"defense"`
			CreatedAt     *string    `db:"created_at"`
		}
	*/

	GetKemonosResponse []GetKemonoResponse

	GetKemonoResponse struct {
		ID            uuid.UUID `json:"id"`
		Image         []byte    `json:"image"`
		Prompt        string    `json:"prompt"`
		Concepts      string    `json:"concepts"`
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
		MaxHp         int       `json:"max_hp"`
		Hp            int       `json:"hp"`
		Attack        int       `json:"attack"`
		Defense       int       `json:"defense"`
		CreatedAt     string    `json:"created_at"`
	}
)

func kemonoToGetKemonoResponse(kemono *repository.Kemono) GetKemonoResponse {
	return GetKemonoResponse{
		ID:            *kemono.ID,
		Image:         kemono.Image,
		Prompt:        *kemono.Prompt,
		Name:          *kemono.Name,
		Description:   *kemono.Description,
		CharacterChip: *kemono.CharacterChip,
		IsPlayer:      *kemono.IsPlayer,
		PlayerID:      *kemono.PlayerID,
		IsOwned:       *kemono.IsOwned,
		OwnerID:       *kemono.OwnerID,
		IsInField:     *kemono.IsInField,
		IsBoss:        *kemono.IsBoss,
		Field:         *kemono.Field,
		X:             *kemono.X,
		Y:             *kemono.Y,
		HasParent:     *kemono.HasParent,
		Parent1ID:     *kemono.Parent1ID,
		Parent2ID:     *kemono.Parent2ID,
		HasChild:      *kemono.HasChild,
		ChildID:       *kemono.ChildID,
		MaxHp:         *kemono.MaxHp,
		Hp:            *kemono.Hp,
		Attack:        *kemono.Attack,
		Defense:       *kemono.Defense,
		CreatedAt:     *kemono.CreatedAt,
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

// GET /api/v1/kemonos/players/:playerID
func (h *Handler) GetKemonoByOwnerId(c echo.Context) error {
	playerID, err := uuid.Parse(c.Param("playerID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid playerID").SetInternal(err)
	}

	kemono, err := h.repo.GetKemonoByOwnerId(c.Request().Context(), playerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := kemonoToGetKemonoResponse(kemono)

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/kemonos
func (h *Handler) CreateKemono(c echo.Context) error {
	id1, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")
	id2, _ := uuid.Parse("00000000-0000-0000-0000-000000000002")
	id3, _ := uuid.Parse("00000000-0000-0000-0000-000000000003")
	id4, _ := uuid.Parse("00000000-0000-0000-0000-000000000004")

	kemonoParams := &repository.KemonoParams{
		ID:            id1,
		Image:         images.TestKemonoImageAqua,
		Prompt:        "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:      "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:          "アクアフローラ",
		Description:   "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		CharacterChip: 1,
		IsPlayer:      true,
		PlayerID:      id1,
		IsOwned:       true,
		OwnerID:       id1,
		IsInField:     false,
		IsBoss:        false,
		Field:         1,
		X:             10,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		MaxHp:         100,
		Hp:            100,
		Attack:        10,
		Defense:       10,
	}
	createdKemonoUUID, err := h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &repository.KemonoParams{
		ID:            id2,
		Image:         images.TestKemonoImageFire,
		Prompt:        "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:      "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,炎属性",
		Name:          "ヒバナ",
		Description:   "このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。",
		CharacterChip: 2,
		IsPlayer:      false,
		PlayerID:      uuid.Nil,
		IsOwned:       true,
		OwnerID:       id1,
		IsInField:     false,
		IsBoss:        false,
		Field:         1,
		X:             10,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		MaxHp:         20,
		Hp:            20,
		Attack:        5,
		Defense:       3,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &repository.KemonoParams{
		ID:            id3,
		Image:         images.TestKemonoImageFire,
		Prompt:        "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:      "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,炎属性",
		Name:          "ヒバナ",
		Description:   "このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。",
		CharacterChip: 3,
		IsPlayer:      false,
		PlayerID:      uuid.Nil,
		IsOwned:       false,
		OwnerID:       uuid.Nil,
		IsInField:     true,
		IsBoss:        false,
		Field:         1,
		X:             13,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		MaxHp:         30,
		Hp:            30,
		Attack:        7,
		Defense:       4,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &repository.KemonoParams{
		ID:            id4,
		Image:         images.TestKemonoImageAqua,
		Prompt:        "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:      "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:          "アクアフローラ",
		Description:   "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		CharacterChip: 1,
		IsPlayer:      false,
		PlayerID:      uuid.Nil,
		IsOwned:       false,
		OwnerID:       uuid.Nil,
		IsInField:     true,
		IsBoss:        false,
		Field:         1,
		X:             16,
		Y:             10,
		HasParent:     false,
		Parent1ID:     uuid.Nil,
		Parent2ID:     uuid.Nil,
		HasChild:      false,
		ChildID:       uuid.Nil,
		MaxHp:         100,
		Hp:            100,
		Attack:        10,
		Defense:       10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdKemonoUUID)
}

// GET /api/v1/kemonos/reset
func (h *Handler) ResetKemonos(c echo.Context) error {
	if err := h.repo.ResetKemonos(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}

// POST /api/v1/kemonos/:kemonoID/catch
func (h *Handler) CatchKemono(c echo.Context) error {
	kemonoID, err := uuid.Parse(c.Param("kemonoID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemonoID").SetInternal(err)
	}

	playerId, err := uuid.Parse(c.FormValue("player_id"))
	if err != nil || playerId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid player_id").SetInternal(err)
	}

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	if *kemono.IsOwned {
		return echo.NewHTTPError(http.StatusConflict, "kemono is already owned")
	}
	if *kemono.IsPlayer {
		return echo.NewHTTPError(http.StatusConflict, "kemono is player")
	}

	t := true
	f := false

	kemono.IsOwned = &t
	kemono.OwnerID = &playerId
	kemono.IsInField = &f

	if err = h.repo.UpdateKemono(c.Request().Context(), kemono); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
