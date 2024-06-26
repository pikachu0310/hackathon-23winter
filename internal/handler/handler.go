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
		userAPI.POST("/login", h.Login)
		userAPI.POST("/signup", h.Signup)
		userAPI.POST("/:userID", h.CreateUserByUserID)
		userAPI.GET("/:userID", h.GetUser)

		userAPI.GET("/:userID/me", h.GetMyKemono)
		userAPI.POST("/:userID/position", h.PostPositionByUser)
		userAPI.GET("/:userID/kemonos", h.GetKemonoByOwnerId)
		userAPI.GET("/:userID/kemonos/me", h.GetMyKemono)
		userAPI.GET("/:userID/kemonos/noimage", h.GetKemonoByOwnerIdNoImage)
		userAPI.GET("/:userID/kemonos/battle", h.GetKemonoForBattleByOwnerId)
		userAPI.POST("/:userID/kemonos/battle", h.PostBattleByPlayerId)
		userAPI.POST("/:userID/kemonos/position", h.PostPositionByUser)
		userAPI.GET("/:userID/concepts", h.GetConceptsByPlayerId)
	}

	// kemono API
	kemonoAPI := api.Group("/kemonos")
	{
		kemonoAPI.GET("", h.GetKemonos)
		kemonoAPI.POST("", h.CreateKemono)
		kemonoAPI.GET("/:kemonoID", h.GetKemono)
		kemonoAPI.GET("/:kemonoID/image", h.GetKemonoImage)
		kemonoAPI.POST("/:kemonoID/catch", h.CatchKemono)
		kemonoAPI.POST("/:kemonoID/extract", h.ExtractKemono)
		kemonoAPI.POST("/generate", h.GenerateKemono)
		kemonoAPI.POST("/breed", h.BreedKemono)
	}

	// field API
	fieldAPI := api.Group("/fields")
	{
		fieldAPI.GET("/:fieldID/kemonos", h.GetKemonosByField)
		fieldAPI.GET("/:fieldID/kemonos/generate", h.GenerateFieldKemono)
		fieldAPI.GET("/:fieldID/kemonos/normal", h.GetNormalKemonoByField)
	}

	// concept API
	conceptAPI := api.Group("/concepts")
	{
		conceptAPI.GET("", h.GetConcepts)
		conceptAPI.GET("/:conceptID", h.GetConcept)
		conceptAPI.DELETE("/:conceptID", h.DeleteConcept)
	}

	// battle API
	battleAPI := api.Group("/battles")
	{
		battleAPI.GET("", h.GetBattles)
		battleAPI.POST("", h.PostBattle)
		battleAPI.GET("/:battle_id", h.GetBattle)
		battleAPI.POST("/:battle_id", h.PostBattleDamage)
	}

	// test API
	testAPI := api.Group("/test")
	{
		testAPI.GET("", h.Test)
	}
	//	testAPI.GET("/2", h.Test2)
	//	testAPI.GET("/3", h.Test3)
	//	testAPI.GET("/4", h.Test4)
	//	testAPI.GET("/5", h.Test5)
	//}

	// reset API
	resetAPI := api.Group("/reset")
	{
		resetAPI.POST("/users", h.ResetUsers)
		resetAPI.POST("/kemonos", h.ResetKemonos)
		resetAPI.POST("/battles", h.ResetBattles)
		resetAPI.POST("/concepts", h.ResetConcepts)
	}
}
