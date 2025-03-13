package api

import "github.com/gin-gonic/gin"

func (server *Server) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  "healthy",
		"message": "Server is running",
	})
}
