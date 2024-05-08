package api

import (
	"github.com/gin-gonic/gin"
)

// GetRouter creates and returns a new HTTP router.
//
// The router is configured with the following routes:
// - /api/v1/items
// - /api/v1/items/:id
// There is a test route at /api/v1/test/notification
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
	// FIXME: This is a test route and should be removed
	setupTestRoutes(v1Group)
}

func setupItemsRoutes(group *gin.RouterGroup) {
	group.GET("/items", getItems)
	group.GET("/items/:id", getItem)
	group.POST("/items", postItem)
	group.PUT("/items/:id", updateItem)
	group.DELETE("/items/:id", deleteItem)
}

// setupTestRoutes sets up the test routes.
// This function is only included as a proof of concept
func setupTestRoutes(group *gin.RouterGroup) {
	group.GET("/test/notification", sendNotification)
}
