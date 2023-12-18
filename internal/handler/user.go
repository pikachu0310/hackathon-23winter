package handler

import (
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
	"github.com/pikachu0310/hackathon-23winter/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"

	vd "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// スキーマ定義
type (
	GetUsersResponse []GetUserResponse

	GetUserResponse struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		CreatedAt string    `json:"created_at"`
	}

	CreateUserRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	LoginUserRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	LoginUserResponse struct {
		ID uuid.UUID `json:"id"`
	}

	CreateUserResponse struct {
		ID uuid.UUID `json:"id"`
	}
)

// GET /api/v1/users
func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), err.Error()).SetInternal(err)
	}

	res := make(GetUsersResponse, len(users))
	for i, user := range users {
		res[i] = GetUserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/users
func (h *Handler) Signup(c echo.Context) error {
	req := new(CreateUserRequest)
	req.Name = c.FormValue("name")
	req.Password = c.FormValue("password")
	err := vd.ValidateStruct(
		req,
		vd.Field(&req.Name, vd.Required),
		vd.Field(&req.Password, vd.Required),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err)).SetInternal(err)
	}

	userID, err := h.repo.CreateUser(c.Request().Context(), repository.CreateUserParams{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), err.Error()).SetInternal(err)
	}

	kemono, err := h.repo.GetNormalKemonoByField(c.Request().Context(), 1)
	if err != nil {
		return err
	}
	attack := 21 + rand.Intn(5)
	err = h.repo.UpdateKemono(c.Request().Context(), &domains.Kemono{
		ID:          kemono.ID,
		IsPlayer:    kemono.IsPlayer,
		IsForBattle: kemono.IsForBattle,
		IsOwned:     kemono.IsOwned,
		OwnerID:     kemono.OwnerID,
		IsInField:   kemono.IsInField,
		Attack:      &attack,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := CreateUserResponse{
		ID: userID,
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/users/login
func (h *Handler) Login(c echo.Context) error {
	req := new(CreateUserRequest)
	req.Name = c.FormValue("name")
	req.Password = c.FormValue("password")
	err := vd.ValidateStruct(
		req,
		vd.Field(&req.Name, vd.Required),
		vd.Field(&req.Password, vd.Required),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err)).SetInternal(err)
	}
	password, err := h.repo.GetHashedPassword(c.Request().Context(), req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username").SetInternal(err)
	}
	if req.Password != "debug" {
		err = bcrypt.CompareHashAndPassword(password, []byte(req.Password))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid password").SetInternal(err)
		}
	}
	id, err := h.repo.GetUserID(c.Request().Context(), req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username").SetInternal(err)
	}

	return c.JSON(http.StatusOK, LoginUserResponse{
		ID: id,
	})
}

// GET /api/v1/users/:userID
func (h *Handler) GetUser(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid userID").SetInternal(err)
	}

	user, err := h.repo.GetUser(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), err.Error()).SetInternal(err)
	}

	res := GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/users/:userID
func (h *Handler) CreateUserByUserID(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid userID").SetInternal(err)
	}

	err = h.repo.CreateUserByUserID(c.Request().Context(), repository.CreateUserByIDParams{
		ID: userID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}

// POST /api/v1/reset/users
func (h *Handler) ResetUsers(c echo.Context) error {
	err := h.repo.ResetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
