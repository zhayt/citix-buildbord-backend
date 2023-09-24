package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetAllNews(e echo.Context) error {
	zap.S().Info("Start getting all news")
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	news, err := h.service.News.GetNews(ctx)
	if err != nil {
		zap.S().Errorf("Getting all news error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("News got successfully")
	return e.JSON(http.StatusOK, news)
}
