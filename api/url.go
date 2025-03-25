package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjthecoder65/url-shortener/db"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ErrInvalidURL        = "invalid url format"
	ErrShortCodeRequired = "short url code is required"
	ErrShortCodeNotFound = "short url not found"
)

type CreateShortURLRequest struct {
	URL string `json:"url" binding:"required,httpurl"`
}

func (server *Server) CreateShortURL(ctx *gin.Context) {
	var req CreateShortURLRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.logger.WithFields(logrus.Fields{
			"request": req,
			"error":   err.Error(),
		}).Error(ErrInvalidURL)
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrInvalidURL))
		return
	}

	arg := db.CreateShortURLParams{
		URL: req.URL,
	}

	dbContext, cancel := context.WithTimeout(context.Background(), server.config.DBTimeout)
	defer cancel()

	url, err := server.queries.CreateShortURL(dbContext, server.config, arg)

	if err != nil {
		server.logger.WithFields(logrus.Fields{
			"url":   req.URL,
			"error": err.Error(),
		}).Error("Failed to create short URL")
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to create short url"))
		return
	}

	server.logger.WithFields(logrus.Fields{
		"url":      req.URL,
		"shortURL": url.ShortCode,
	}).Info("Short URL Created successfully")

	ctx.JSON(http.StatusCreated, url)
}

func (server *Server) GetOriginalURL(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")

	if shortCode == "" {
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
		}).Warn(ErrShortCodeRequired)
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrShortCodeRequired))
		return
	}

	dbContext, cancel := context.WithTimeout(context.Background(), server.config.DBTimeout)
	defer cancel()

	url, err := server.queries.GetShortURL(dbContext, shortCode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			server.logger.WithFields(logrus.Fields{
				"shortCode": shortCode,
			}).Warn(ErrShortCodeNotFound)
			ctx.JSON(http.StatusNotFound, errorResponse(ErrShortCodeNotFound))
			return
		}

		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
			"error":     err.Error(),
		}).Error("Failed to fetch short URL")
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to fetch short url"))
		return
	}

	server.logger.WithFields(logrus.Fields{
		"shortCode": shortCode,
		"url":       url.URL,
	}).Info("Short URL retried successfully")

	ctx.JSON(http.StatusOK, url)
}

type UpdateShortURLRequest struct {
	URL string `json:"url" binding:"required,httpurl"`
}

func (server *Server) UpdateShortURL(ctx *gin.Context) {
	var req UpdateShortURLRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.logger.WithFields(logrus.Fields{
			"request": req,
			"error":   err.Error(),
		}).Error(ErrInvalidURL)
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrInvalidURL))
		return
	}

	shortCode := ctx.Param("shortCode")

	if shortCode == "" {
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
		}).Warn(ErrShortCodeRequired)
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrShortCodeRequired))
		return
	}

	arg := db.UpdateShortURLParams{
		ShortCode: shortCode,
		URL:       req.URL,
	}

	dbContext, cancel := context.WithTimeout(context.Background(), server.config.DBTimeout)
	defer cancel()

	url, err := server.queries.UpdateShortURL(dbContext, arg)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			server.logger.WithFields(logrus.Fields{
				"shortCode": shortCode,
			}).Warn(ErrShortCodeNotFound)
			ctx.JSON(http.StatusNotFound, errorResponse(ErrShortCodeNotFound))
			return
		}
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
			"error":     err.Error(),
		}).Error("Failed to update short URL")
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to update short url"))
		return
	}

	server.logger.WithFields(logrus.Fields{
		"shortCode": shortCode,
		"url":       req.URL,
	}).Info("Short URL updated successfully")

	ctx.JSON(http.StatusOK, url)
}

func (server *Server) DeleteShortURL(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")

	if shortCode == "" {
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
		}).Warn(ErrShortCodeRequired)
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrShortCodeRequired))
		return
	}

	dbContext, cancel := context.WithTimeout(context.Background(), server.config.DBTimeout)
	defer cancel()

	err := server.queries.DeleteShortURL(dbContext, shortCode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			server.logger.WithFields(logrus.Fields{
				"shortCode": shortCode,
			}).Warn(ErrShortCodeNotFound)
			ctx.JSON(http.StatusNotFound, errorResponse(ErrShortCodeNotFound))
			return
		}
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
			"error":     err.Error(),
		}).Error("Failed to delete short URL")
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to delete short url"))
		return
	}

	server.logger.WithFields(logrus.Fields{
		"shortCode": shortCode,
	}).Info("Short URL delete successfully")

	ctx.Status(http.StatusNoContent)
}

func (server *Server) GetURLStats(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")

	if shortCode == "" {
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
		}).Warn(ErrShortCodeRequired)
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrShortCodeRequired))
		return
	}

	dbContext, cancel := context.WithTimeout(context.Background(), server.config.DBTimeout)
	defer cancel()

	url, err := server.queries.GetShortURLStats(dbContext, shortCode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			server.logger.WithFields(logrus.Fields{
				"shortCode": shortCode,
			}).Warn(ErrShortCodeNotFound)
			ctx.JSON(http.StatusNotFound, errorResponse(ErrShortCodeNotFound))
			return
		}
		server.logger.WithFields(logrus.Fields{
			"shortCode": shortCode,
			"error":     err.Error(),
		}).Error("Failed to Fetch short URL")
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to fetch url"))
		return
	}

	server.logger.WithFields(logrus.Fields{
		"shortCode": shortCode,
	}).Info("Short URL stats retrieved succefully")

	ctx.JSON(http.StatusOK, url)
}

func errorResponse(message string) gin.H {
	return gin.H{"error": message}
}
