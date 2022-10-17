package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createUser)
			lists.GET("/", h.getAllUsers)
			lists.GET("/:id", h.getUserById)
			lists.PUT("/:id", h.updateUser)
			lists.DELETE("/:id", h.deleteUser)
		}
	}
	return router
}
