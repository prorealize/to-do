package api

import (
	"github.com/gin-gonic/gin"
)

// GetRouter creates and returns a new HTTP router.
func GetRouter() *gin.Engine {
	router := gin.Default()
	setupApiRoutes(router)
	return router
}

func setupApiRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	setupV1Routes(apiGroup)
}

func setupV1Routes(group *gin.RouterGroup) {
	v1Group := group.Group("/v1")
	setupItemsRoutes(v1Group)
	setupTestNotificationRoutes(v1Group)
}

func setupItemsRoutes(group *gin.RouterGroup) {
	group.GET("/items", getItems)
	group.GET("/items/:id", getItem)
	group.POST("/items", postItem)
	group.PUT("/items/:id", updateItem)
	group.DELETE("/items/:id", deleteItem)
}

func setupTestNotificationRoutes(group *gin.RouterGroup) {
	group.GET("/test-notification", sendNotification)
}
