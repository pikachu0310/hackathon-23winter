package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/api"
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
	"github.com/pikachu0310/hackathon-23winter/src/images"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
	"time"
)

// スキーマ定義
type (
	/*
		Kemono struct {
			ID            *uuid.UUID    `db:"id"`
			Image         []byte        `db:"image"`
			Prompt        *string       `db:"prompt"`
			Concepts      *ConceptsText `db:"concepts"`
			Name          *string       `db:"name"`
			Description   *string       `db:"description"`
			Kind 		  *int          `db:"kind"`
			Color 		  *int          `db:"color"`
			IsPlayer      *bool         `db:"is_player"`
			IsForBattle   *bool         `db:"is_for_battle"`
			IsOwned       *bool         `db:"is_owned"`
			OwnerID       *uuid.UUID    `db:"owner_id"`
			IsInField     *bool         `db:"is_in_field"`
			IsBoss        *bool         `db:"is_boss"`
			Field         *int          `db:"field"`
			X             *int          `db:"x"`
			Y             *int          `db:"y"`
			HasParent     *bool         `db:"has_parent"`
			Parent1ID     *uuid.UUID    `db:"parent1_id"`
			Parent2ID     *uuid.UUID    `db:"parent2_id"`
			HasChild      *bool         `db:"has_child"`
			ChildID       *uuid.UUID    `db:"child_id"`
			MaxHp         *int          `db:"max_hp"`
			Hp            *int          `db:"hp"`
			Attack        *int          `db:"attack"`
			Defense       *int          `db:"defense"`
			CreatedAt     *string       `db:"created_at"`
		}
	*/

	GetKemonosResponse []GetKemonoResponse

	GetKemonoResponse struct {
		ID          uuid.UUID `json:"id"`
		Image       []byte    `json:"image"`
		Prompt      string    `json:"prompt"`
		Concepts    []string  `json:"concepts"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Kind        int       `json:"kind"`
		Color       int       `json:"color"`
		IsPlayer    bool      `json:"is_player"`
		IsForBattle bool      `json:"is_for_battle"`
		IsOwned     bool      `json:"is_owned"`
		OwnerID     uuid.UUID `json:"owner_id"`
		IsInField   bool      `json:"is_in_field"`
		IsBoss      bool      `json:"is_boss"`
		Field       int       `json:"field"`
		X           int       `json:"x"`
		Y           int       `json:"y"`
		HasParent   bool      `json:"has_parent"`
		Parent1ID   uuid.UUID `json:"parent1_id"`
		Parent2ID   uuid.UUID `json:"parent2_id"`
		HasChild    bool      `json:"has_child"`
		ChildID     uuid.UUID `json:"child_id"`
		MaxHp       int       `json:"max_hp"`
		Hp          int       `json:"hp"`
		Attack      int       `json:"attack"`
		Defense     int       `json:"defense"`
		CreatedAt   string    `json:"created_at"`
	}

	GetKemonosResponseWithoutImage []GetKemonoResponseWithoutImage
	GetKemonoResponseWithoutImage  struct {
		ID          uuid.UUID `json:"id"`
		Prompt      string    `json:"prompt"`
		Concepts    []string  `json:"concepts"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Kind        int       `json:"kind"`
		Color       int       `json:"color"`
		IsPlayer    bool      `json:"is_player"`
		IsForBattle bool      `json:"is_for_battle"`
		IsOwned     bool      `json:"is_owned"`
		OwnerID     uuid.UUID `json:"owner_id"`
		IsInField   bool      `json:"is_in_field"`
		IsBoss      bool      `json:"is_boss"`
		Field       int       `json:"field"`
		X           int       `json:"x"`
		Y           int       `json:"y"`
		HasParent   bool      `json:"has_parent"`
		Parent1ID   uuid.UUID `json:"parent1_id"`
		Parent2ID   uuid.UUID `json:"parent2_id"`
		HasChild    bool      `json:"has_child"`
		ChildID     uuid.UUID `json:"child_id"`
		MaxHp       int       `json:"max_hp"`
		Hp          int       `json:"hp"`
		Attack      int       `json:"attack"`
		Defense     int       `json:"defense"`
		CreatedAt   string    `json:"created_at"`
	}
)

func kemonoToGetKemonoResponse(kemono *domains.Kemono) GetKemonoResponse {
	return GetKemonoResponse{
		ID:          *kemono.ID,
		Image:       kemono.Image,
		Prompt:      *kemono.Prompt,
		Concepts:    kemono.Concepts.Concepts(),
		Name:        *kemono.Name,
		Description: *kemono.Description,
		Kind:        *kemono.Kind,
		Color:       *kemono.Color,
		IsPlayer:    *kemono.IsPlayer,
		IsForBattle: *kemono.IsForBattle,
		IsOwned:     *kemono.IsOwned,
		OwnerID:     *kemono.OwnerID,
		IsInField:   *kemono.IsInField,
		IsBoss:      *kemono.IsBoss,
		Field:       *kemono.Field,
		X:           *kemono.X,
		Y:           *kemono.Y,
		HasParent:   *kemono.HasParent,
		Parent1ID:   *kemono.Parent1ID,
		Parent2ID:   *kemono.Parent2ID,
		HasChild:    *kemono.HasChild,
		ChildID:     *kemono.ChildID,
		MaxHp:       *kemono.MaxHp,
		Hp:          *kemono.Hp,
		Attack:      *kemono.Attack,
		Defense:     *kemono.Defense,
		CreatedAt:   *kemono.CreatedAt,
	}
}

func (kemono GetKemonoResponse) WithoutImage() GetKemonoResponseWithoutImage {
	return GetKemonoResponseWithoutImage{
		ID:          kemono.ID,
		Prompt:      kemono.Prompt,
		Concepts:    kemono.Concepts,
		Name:        kemono.Name,
		Description: kemono.Description,
		Kind:        kemono.Kind,
		Color:       kemono.Color,
		IsPlayer:    kemono.IsPlayer,
		IsForBattle: kemono.IsForBattle,
		IsOwned:     kemono.IsOwned,
		OwnerID:     kemono.OwnerID,
		IsInField:   kemono.IsInField,
		IsBoss:      kemono.IsBoss,
		Field:       kemono.Field,
		X:           kemono.X,
		Y:           kemono.Y,
		HasParent:   kemono.HasParent,
		Parent1ID:   kemono.Parent1ID,
		Parent2ID:   kemono.Parent2ID,
		HasChild:    kemono.HasChild,
		ChildID:     kemono.ChildID,
		MaxHp:       kemono.MaxHp,
		Hp:          kemono.Hp,
		Attack:      kemono.Attack,
		Defense:     kemono.Defense,
		CreatedAt:   kemono.CreatedAt,
	}
}

func (kemonos GetKemonosResponse) WithoutImage() GetKemonosResponseWithoutImage {
	res := make(GetKemonosResponseWithoutImage, len(kemonos))
	for i, kemono := range kemonos {
		res[i] = kemono.WithoutImage()
	}
	return res
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
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "kemono not found").SetInternal(err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := kemonoToGetKemonoResponse(kemono)

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/kemonos/:kemonoID/image
func (h *Handler) GetKemonoImage(c echo.Context) error {
	kemonoID, err := uuid.Parse(c.Param("kemonoID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemonoID").SetInternal(err)
	}

	kemonoImage, err := h.repo.GetKemonoImage(c.Request().Context(), kemonoID)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "kemono not found").SetInternal(err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.Blob(http.StatusOK, "image/png", kemonoImage)

}

// GET /api/v1/fields/:fieldID/kemonos
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

	return c.JSON(http.StatusOK, res.WithoutImage())
}

// GET /api/v1/users/:userID/kemonos
func (h *Handler) GetKemonoByOwnerId(c echo.Context) error {
	playerID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid playerID").SetInternal(err)
	}

	kemonos, err := h.repo.GetKemonoByOwnerId(c.Request().Context(), playerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoPlayerFound := false
	kemonoPlayer := domains.Kemono{}
	kemonoBattlerFound := false
	kemonoBattler := domains.Kemono{}
	for _, kemono := range kemonos {
		if *kemono.IsPlayer {
			kemonoPlayer = kemono
			kemonoPlayerFound = true
		}
		if *kemono.IsForBattle {
			kemonoBattler = kemono
			kemonoBattlerFound = true
		}
	}
	res := GetKemonosResponse{}
	if kemonoPlayerFound {
		res = append(res, kemonoToGetKemonoResponse(&kemonoPlayer))
	}
	if kemonoBattlerFound {
		res = append(res, kemonoToGetKemonoResponse(&kemonoBattler))
	}
	for _, kemono := range kemonos {
		if kemonoPlayerFound {
			if *kemono.ID == *kemonoPlayer.ID {
				continue
			}
		}
		if kemonoBattlerFound {
			if *kemono.ID == *kemonoBattler.ID {
				continue
			}
		}
		res = append(res, kemonoToGetKemonoResponse(&kemono))
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/users/:userID/me
func (h *Handler) GetMyKemono(c echo.Context) error {
	playerID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid playerID").SetInternal(err)
	}

	kemono, err := h.repo.GetMyKemonoByUserId(c.Request().Context(), playerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := kemonoToGetKemonoResponse(kemono)

	return c.JSON(http.StatusOK, res.WithoutImage())
}

// GET /api/v1/users/:userID/kemonos/battle
func (h *Handler) GetKemonoForBattleByOwnerId(c echo.Context) error {
	playerID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid playerID").SetInternal(err)
	}

	kemono, err := h.repo.GetKemonoForBattleByOwnerId(c.Request().Context(), playerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := kemonoToGetKemonoResponse(kemono)

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/users/:userID/kemonos/battle
func (h *Handler) PostBattleByPlayerId(c echo.Context) error {
	playerID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid playerID").SetInternal(err)
	}

	newKemonoForBattleId, err := uuid.Parse(c.FormValue("kemono_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemono_id").SetInternal(err)
	}
	newKemono, err := h.repo.GetKemono(c.Request().Context(), newKemonoForBattleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	oldKemono, err := h.repo.GetKemonoForBattleByOwnerId(c.Request().Context(), playerID)
	if errors.Is(err, sql.ErrNoRows) {
		err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
			ID:          newKemono.ID,
			IsForBattle: domains.NewBool(true),
		})
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusOK)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	if *newKemono.OwnerID != playerID {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid kemono_owner_id:%s user_id:%s", *newKemono.OwnerID, playerID)).SetInternal(err)
	}

	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:          newKemono.ID,
		IsForBattle: domains.NewBool(true),
	})
	if err != nil {
		return err
	}
	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:          oldKemono.ID,
		IsForBattle: domains.NewBool(false),
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// POST /api/v1/kemonos
func (h *Handler) CreateKemono(c echo.Context) error {
	id1, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")
	id2, _ := uuid.Parse("00000000-0000-0000-0000-000000000002")
	id3, _ := uuid.Parse("00000000-0000-0000-0000-000000000003")
	id4, _ := uuid.Parse("00000000-0000-0000-0000-000000000004")

	kemonoParams := &domains.KemonoParams{
		ID:          id1,
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        0,
		Color:       8,
		IsPlayer:    true,
		IsForBattle: true,
		IsOwned:     true,
		OwnerID:     id1,
		IsInField:   false,
		IsBoss:      false,
		Field:       1,
		X:           10,
		Y:           15,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err := h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          id2,
		Image:       images.TestKemonoImageFire,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,炎属性",
		Name:        "ヒバナ",
		Description: "このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。",
		Kind:        0,
		Color:       3,
		IsPlayer:    false,
		IsOwned:     true,
		OwnerID:     id1,
		IsInField:   false,
		IsBoss:      false,
		Field:       1,
		X:           10,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       20,
		Hp:          20,
		Attack:      5,
		Defense:     3,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          id3,
		Image:       images.TestKemonoImageFire,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,炎属性",
		Name:        "ヒバナ",
		Description: "このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。",
		Kind:        0,
		Color:       3,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           13,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       30,
		Hp:          30,
		Attack:      7,
		Defense:     4,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          id4,
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        0,
		Color:       8,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           16,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        1,
		Color:       3,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           19,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        2,
		Color:       7,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           21,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        3,
		Color:       4,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           13,
		Y:           13,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        4,
		Color:       9,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           17,
		Y:           14,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        5,
		Color:       6,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           27,
		Y:           14,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        6,
		Color:       10,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           25,
		Y:           17,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        7,
		Color:       7,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           27,
		Y:           21,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        1,
		Color:       4,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           19,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        2,
		Color:       8,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           21,
		Y:           10,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        3,
		Color:       5,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           13,
		Y:           13,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        4,
		Color:       10,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           17,
		Y:           14,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        5,
		Color:       7,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           27,
		Y:           14,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        6,
		Color:       11,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           25,
		Y:           17,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	kemonoParams = &domains.KemonoParams{
		ID:          uuid.New(),
		Image:       images.TestKemonoImageAqua,
		Prompt:      "Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.",
		Concepts:    "とてもかわいい,マスコット,四足歩行,目が覚めたら森の中だった,ポケモンのようなイメージ,色は鮮やかめ,水属性",
		Name:        "アクアフローラ",
		Description: "このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。",
		Kind:        7,
		Color:       8,
		IsPlayer:    false,
		IsOwned:     false,
		OwnerID:     uuid.Nil,
		IsInField:   true,
		IsBoss:      false,
		Field:       1,
		X:           27,
		Y:           21,
		HasParent:   false,
		Parent1ID:   uuid.Nil,
		Parent2ID:   uuid.Nil,
		HasChild:    false,
		ChildID:     uuid.Nil,
		MaxHp:       100,
		Hp:          100,
		Attack:      10,
		Defense:     10,
	}
	createdKemonoUUID, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdKemonoUUID)
}

// GET /api/v1/reset/kemonos
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

	if err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:        &kemonoID,
		IsOwned:   domains.NewBool(true),
		OwnerID:   &playerId,
		IsInField: domains.NewBool(false),
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}

// POST /api/v1/kemonos/:kemonoID/extract
func (h *Handler) ExtractKemono(c echo.Context) error {
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

	for _, concept := range kemono.Concepts.Concepts() {
		_, err = h.repo.CreateConcept(playerId, concept)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
	}

	return c.NoContent(http.StatusOK)
}

// POST /api/v1/kemonos/generate
func (h *Handler) GenerateKemono(c echo.Context) error {
	startTime := time.Now()

	userID, err := uuid.Parse(c.FormValue("user_id"))
	if err != nil || userID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user_id").SetInternal(err)
	}
	conceptsString := c.FormValue("concepts")
	if conceptsString == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid concepts")
	}
	conceptsText := domains.ConceptsText(conceptsString)

	kemonoID, err := h.repo.CreateKemono(c.Request().Context(), &domains.Kemono{
		ID:          nil,
		Image:       nil,
		Prompt:      nil,
		Concepts:    &conceptsText,
		Name:        nil,
		Description: nil,
		Kind:        nil,
		Color:       nil,
		IsPlayer:    domains.NewBool(false),
		IsForBattle: domains.NewBool(false),
		IsOwned:     domains.NewBool(true),
		OwnerID:     &userID,
		IsInField:   domains.NewBool(false),
		IsBoss:      domains.NewBool(false),
		Field:       nil,
		X:           nil,
		Y:           nil,
		HasParent:   domains.NewBool(false),
		Parent1ID:   domains.UuidNil(),
		Parent2ID:   domains.UuidNil(),
		HasChild:    domains.NewBool(false),
		ChildID:     domains.UuidNil(),
		MaxHp:       nil,
		Hp:          nil,
		Attack:      nil,
		Defense:     nil,
		CreatedAt:   nil,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	// 4s 3s 4s
	err = h.generateKemonoPromptAndUpdateKemono(c, kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 13s 25s 15s
	err = h.generateKemonoImageAndUpdateKemono(c, kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 16s 15s 31s
	err = h.generateKemonoDescriptionAndUpdateKemono(c, kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 11s 25s
	var eg errgroup.Group
	eg.Go(func() error {
		// 5s
		err = h.generateKemonoStatusAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})

	eg.Go(func() error {
		// 17s
		err = h.generateKemonoCharacterChipAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})

	eg.Go(func() error {
		// 7s
		err = h.generateKemonoNameAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	fmt.Println(time.Now().Sub(startTime).Seconds())

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, kemonoToGetKemonoResponse(kemono))
}

// POST /api/v1/kemonos/breed
func (h *Handler) BreedKemono(c echo.Context) error {
	startTime := time.Now()

	kemono1ID, err := uuid.Parse(c.FormValue("kemono1_id"))
	if err != nil || kemono1ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemono1_id").SetInternal(err)
	}
	kemono2ID, err := uuid.Parse(c.FormValue("kemono2_id"))
	if err != nil || kemono2ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemono2_id").SetInternal(err)
	}

	kemonoParent1, err := h.repo.GetKemono(c.Request().Context(), kemono1ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent2, err := h.repo.GetKemono(c.Request().Context(), kemono2ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	if *kemonoParent1.ChildID != uuid.Nil || *kemonoParent2.ChildID != uuid.Nil {
		return echo.NewHTTPError(http.StatusConflict, "kemono is already bred")
	}

	kemonoID, err := h.repo.CreateKemono(c.Request().Context(), &domains.Kemono{
		ID:          nil,
		Image:       nil,
		Prompt:      nil,
		Concepts:    nil,
		Name:        nil,
		Description: nil,
		Kind:        nil,
		Color:       nil,
		IsPlayer:    domains.NewBool(false),
		IsForBattle: domains.NewBool(false),
		IsOwned:     domains.NewBool(false),
		OwnerID:     domains.UuidNil(),
		IsInField:   domains.NewBool(false),
		IsBoss:      domains.NewBool(false),
		Field:       nil,
		X:           nil,
		Y:           nil,
		HasParent:   domains.NewBool(true),
		Parent1ID:   kemonoParent1.ID,
		Parent2ID:   kemonoParent2.ID,
		HasChild:    domains.NewBool(false),
		ChildID:     domains.UuidNil(),
		MaxHp:       nil,
		Hp:          nil,
		Attack:      nil,
		Defense:     nil,
		CreatedAt:   nil,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:      kemonoParent1.ID,
		ChildID: &kemonoID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:      kemonoParent2.ID,
		ChildID: &kemonoID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	// 12s
	err = h.generateBreedKemonoPromptAndUpdateKemono(c, kemonoID)
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 15s
	err = h.generateBreedKemonoImageAndUpdateKemono(c, kemonoID)
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 11s
	err = h.generateBreedKemonoDescriptionAndUpdateKemono(c, kemonoID)
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 8s
	err = h.generateBreedKemonoConceptsAndUpdateKemono(c, kemonoID)
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	err = h.generateBreedKemonoStatusAndUpdateKemono(c, kemonoID)
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	var eg errgroup.Group
	eg.Go(func() error {
		err = h.generateKemonoCharacterChipAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})
	eg.Go(func() error {
		err = h.generateKemonoNameAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})
	err = eg.Wait()
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		h.revertChildID(c, kemonoParent1, kemonoParent2)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoParent1, err = h.repo.GetKemono(c.Request().Context(), kemono1ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	kemonoParent2, err = h.repo.GetKemono(c.Request().Context(), kemono2ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	isPlayer := *kemonoParent1.IsPlayer || *kemonoParent2.IsPlayer
	isForBattle := *kemonoParent1.IsForBattle || *kemonoParent2.IsForBattle
	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:          &kemonoID,
		OwnerID:     kemonoParent1.OwnerID,
		IsOwned:     domains.NewBool(true),
		IsForBattle: &isForBattle,
		IsPlayer:    &isPlayer,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:       kemonoParent1.ID,
		HasChild: domains.NewBool(true),
		ChildID:  &kemonoID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:       kemonoParent2.ID,
		HasChild: domains.NewBool(true),
		ChildID:  &kemonoID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, kemonoToGetKemonoResponse(kemono))
}

func (h *Handler) revertChildID(c echo.Context, kemonoParent1 *domains.Kemono, kemonoParent2 *domains.Kemono) {
	_ = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:      kemonoParent1.ID,
		ChildID: &uuid.Nil,
	})
	_ = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:      kemonoParent2.ID,
		ChildID: &uuid.Nil,
	})
	return
}

// POST /api/v1/fields/:fieldID/kemonos/generate
func (h *Handler) GenerateFieldKemono(c echo.Context) error {
	startTime := time.Now()

	fieldID, err := strconv.Atoi(c.Param("fieldID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid fieldID").SetInternal(err)
	}

	fieldType := domains.FieldIdToFieldType(fieldID)
	x, y := domains.FieldTypeToPosition(fieldType)
	concepts := domains.Concepts{}
	kemonoConcepts, err := api.GenerateKemonoConcepts(fieldID)
	if err != nil {
		return err
	}
	for _, concept := range *kemonoConcepts {
		concepts.Add(concept)
	}
	concepts.Add(domains.FieldTypeToConcepts1(fieldType).SelectConceptByRandom())
	concepts.Add(domains.FieldTypeToConcepts2(fieldType).SelectConceptByRandom())

	kemonoID, err := h.repo.CreateKemono(c.Request().Context(), &domains.Kemono{
		ID:          nil,
		Image:       nil,
		Prompt:      nil,
		Concepts:    concepts.Text(),
		Name:        nil,
		Description: nil,
		Kind:        nil,
		Color:       nil,
		IsPlayer:    domains.NewBool(false),
		IsForBattle: domains.NewBool(false),
		IsOwned:     domains.NewBool(false),
		OwnerID:     domains.UuidNil(),
		IsInField:   domains.NewBool(true),
		IsBoss:      domains.NewBool(false),
		Field:       &fieldID,
		X:           &x,
		Y:           &y,
		HasParent:   domains.NewBool(false),
		Parent1ID:   domains.UuidNil(),
		Parent2ID:   domains.UuidNil(),
		HasChild:    domains.NewBool(false),
		ChildID:     domains.UuidNil(),
		MaxHp:       nil,
		Hp:          nil,
		Attack:      nil,
		Defense:     nil,
		CreatedAt:   nil,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	// 4s 3s 4s
	err = h.generateKemonoPromptAndUpdateKemono(c, kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 13s 25s 15s
	err = h.generateKemonoImageAndUpdateKemono(c, kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 16s 15s 31s
	err = h.generateKemonoDescriptionAndUpdateKemono(c, kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	fmt.Println(time.Now().Sub(startTime).Seconds())

	// 11s 25s
	var eg errgroup.Group
	eg.Go(func() error {
		// 5s
		err = h.generateKemonoStatusAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})

	eg.Go(func() error {
		// 17s
		err = h.generateKemonoCharacterChipAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})

	eg.Go(func() error {
		// 7s
		err = h.generateKemonoNameAndUpdateKemono(c, kemonoID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	fmt.Println(time.Now().Sub(startTime).Seconds())

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, kemonoToGetKemonoResponse(kemono))
}

// POST /api/v1/fields/:fieldID/kemonos/normal
func (h *Handler) GetNormalKemonoByField(c echo.Context) error {
	normalKemono, err := h.repo.GetNormalKemonoByField(c.Request().Context(), 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, kemonoToGetKemonoResponse(normalKemono))
}

// POST /api/v1/users/:userID/positions
func (h *Handler) PostPositionByUser(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user_id").SetInternal(err)
	}
	fieldID, err := strconv.Atoi(c.FormValue("field_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid field_id").SetInternal(err)
	}
	x, err := strconv.Atoi(c.FormValue("x"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid player_id").SetInternal(err)
	}
	y, err := strconv.Atoi(c.FormValue("y"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid player_id").SetInternal(err)
	}

	myKemono, err := h.repo.GetMyKemonoByUserId(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:        myKemono.ID,
		Field:     &fieldID,
		IsInField: domains.NewBool(true),
		X:         &x,
		Y:         &y,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
