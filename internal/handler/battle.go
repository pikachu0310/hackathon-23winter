package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/api"
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
	"net/http"
	"strconv"
)

type (
	/*
		CREATE TABLE IF NOT EXISTS battle (
		    id CHAR(36) NOT NULL,
		    my_kemono_id CHAR(36) NOT NULL,
		    enemy_kemono_id CHAR(36) NOT NULL,
		    text TEXT DEFAULT '',
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    PRIMARY KEY (id)
		);
	*/

	PostBattleRequest struct {
		MyKemonoID    uuid.UUID `json:"my_kemono_id"`
		EnemyKemonoID uuid.UUID `json:"enemy_kemono_id"`
	}

	PostBattleResponse struct {
		BattleId uuid.UUID `json:"battle_id"`
	}

	PostBattleDamageRequest struct {
		Damage int `json:"damage"`
	}

	PostBattleDamageResponse struct {
		Text string `json:"text"`
	}
)

// GET /api/v1/battles
func (h *Handler) GetBattles(c echo.Context) error {
	battles, err := h.repo.GetBattles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, battles)
}

// GET /api/v1/battles/:battle_id
func (h *Handler) GetBattle(c echo.Context) error {
	battleId, err := uuid.Parse(c.Param("battle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid battle id").SetInternal(err)
	}

	battle, err := h.repo.GetBattle(battleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, battle)
}

// POST /api/v1/battles
func (h *Handler) PostBattle(c echo.Context) error {
	req := new(PostBattleRequest)
	//if err := c.Bind(req); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	//}
	myKemonoID, err := uuid.Parse(c.FormValue("my_kemono_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid my kemono id").SetInternal(err)
	}
	enemyKemonoID, err := uuid.Parse(c.FormValue("enemy_kemono_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid enemy kemono id").SetInternal(err)
	}
	req.MyKemonoID = myKemonoID
	req.EnemyKemonoID = enemyKemonoID
	if req.MyKemonoID == uuid.Nil || req.EnemyKemonoID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	//if err := c.Validate(req); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	//}

	myKemono, err := h.repo.GetKemono(c.Request().Context(), req.MyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	enemyKemono, err := h.repo.GetKemono(c.Request().Context(), req.EnemyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	battleId, err := h.repo.CreateBattle(req.MyKemonoID, req.EnemyKemonoID, DefaultBattleText(*myKemono, *enemyKemono))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := PostBattleResponse{
		BattleId: battleId,
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/battles/:battle_id
func (h *Handler) PostBattleDamage(c echo.Context) error {
	battleId, err := uuid.Parse(c.Param("battle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid battle id").SetInternal(err)
	}

	req := new(PostBattleDamageRequest)
	//if err := c.Bind(req); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	//}
	damage, err := strconv.Atoi(c.FormValue("damage"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid damage").SetInternal(err)
	}
	req.Damage = damage

	battle, err := h.repo.GetBattle(battleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	myKemono, err := h.repo.GetKemono(c.Request().Context(), *battle.MyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	enemyKemono, err := h.repo.GetKemono(c.Request().Context(), *battle.EnemyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	battleText, err := api.GenerateBattleText(myKemono, enemyKemono, req.Damage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	//text, err := api.GetGeneratedText(fmt.Sprintf(prompt))
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	//}

	err = h.repo.UpdateBattleText(battleId, *battle.Text+"\n"+*battleText)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := PostBattleDamageResponse{
		Text: *battleText,
	}

	return c.JSON(http.StatusOK, res)
}

func DefaultBattleText(attacker domains.Kemono, defender domains.Kemono) string {
	return fmt.Sprintf("")
}

// GET /api/v1/reset/battles
func (h *Handler) ResetBattles(c echo.Context) error {
	err := h.repo.ResetBattles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
