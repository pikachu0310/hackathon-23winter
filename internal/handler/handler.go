package handler

import (
	"github.com/pikachu0310/hackathon-23winter/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) SetupRoutes(api *echo.Group) {
	// ping API
	pingAPI := api.Group("/ping")
	{
		pingAPI.GET("", h.Ping)
	}

	// user API
	userAPI := api.Group("/users")
	{
		userAPI.GET("", h.GetUsers)
		userAPI.POST("", h.CreateUser)
		userAPI.GET("/:userID", h.GetUser)
	}

	// kemono API
	kemonoAPI := api.Group("/kemonos")
	{
		kemonoAPI.GET("", h.GetKemonos)
		kemonoAPI.POST("", h.CreateKemono)
		kemonoAPI.GET("/:kemonoID", h.GetKemono)
		kemonoAPI.POST("/:kemonoID/catch", h.CatchKemono)
		kemonoAPI.GET("/fields/:fieldID", h.GetKemonosByField)
		kemonoAPI.GET("/reset", h.ResetKemonos)
	}

	// battle API
	battleAPI := api.Group("/battles")
	{
		battleAPI.GET("", h.GetBattles)
		battleAPI.POST("", h.PostBattle)
		battleAPI.GET("/:battle_id", h.GetBattle)
		battleAPI.POST("/:battle_id", h.PostBattleDamage)
		battleAPI.GET("/reset", h.ResetBattles)
	}

	// test API
	testAPI := api.Group("/test")
	{
		testAPI.GET("", h.Test)
	}
}
