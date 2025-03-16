package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjthecoder65/url-shortener/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateShortURLRequest struct {
	URL string `json:"url" binding:"required,httpurl"`
}

func (server *Server) CreateShortURL(ctx *gin.Context) {
	var req CreateShortURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateShortURLParams{
		URL: req.URL,
	}

	url, err := server.queries.CreateShortURL(context.Background(), server.config, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, url)
}

func (server *Server) GetOriginalURL(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")
	url, err := server.queries.GetShortURL(context.Background(), shortCode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, url)
}

type UpdateShortURLRequest struct {
	URL string `json:"url" binding:"required,httpurl"`
}

func (server *Server) UpdateShortURL(ctx *gin.Context) {
	var req UpdateShortURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	shortCode := ctx.Param("shortCode")

	arg := db.UpdateShortURLParams{
		ShortCode: shortCode,
		URL:       req.URL,
	}

	url, err := server.queries.UpdateShortURL(context.Background(), arg)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, url)
}

func (server *Server) DeleteShortURL(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")

	err := server.queries.DeleteShortURL(context.Background(), shortCode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (server *Server) GetURLStats(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")

	url, err := server.queries.GetShortURLStats(context.Background(), shortCode)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, url)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err}
}
