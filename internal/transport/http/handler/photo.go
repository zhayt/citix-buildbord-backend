package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) UploadPhoto(e echo.Context) error {
	zap.S().Info("Start upload photo")
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	zap.S().Info("Read request body")
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		zap.S().Errorf("Reading request error: %s", err.Error())
		return e.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	//zap.S().Info("Convert from base64")
	//imageBinary, err := base64.StdEncoding.DecodeString(string(body))
	//if err != nil {
	//	zap.S().Errorf("Converting base64 error: %s", err.Error())
	//	return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	//}

	zap.S().Info("Start save photo")
	if err = h.service.Photo.SavePhoto(ctx, string(body)); err != nil {
		zap.S().Errorf("Saving photo error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("Photo saved successfully")
	return e.NoContent(http.StatusOK)
}

func (h *Handler) DownloadPhoto(e echo.Context) error {
	zap.S().Info("Start download photo")
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	photoInfo, err := h.service.Photo.GetAllPhoto(ctx)
	if err != nil {
		zap.S().Errorf("Getting photo error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("All photo send successfully")
	return e.JSON(http.StatusOK, photoInfo)
}
