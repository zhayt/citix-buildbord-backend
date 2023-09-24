package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"innovatex-app/internal/models"
	"net/http"
)

func (h *Handler) GetSurveyList(e echo.Context) error {
	zap.S().Info("Start getting survey list")
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	surveyInfo, err := h.service.Survey.GetSurveyList(ctx)
	if err != nil {
		zap.S().Errorf("Getting survey list error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("All survey list got successfully")
	return e.JSON(http.StatusOK, surveyInfo)
}

func (h *Handler) GetSurvey(e echo.Context) error {
	zap.S().Info("Start getting survey by id")
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	surveyID := e.Param("survey_id")
	survey, err := h.service.Survey.GetSurvey(ctx, surveyID)
	if err != nil {
		zap.S().Errorf("Getting servey error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	if survey == nil {
		zap.S().Errorf("Survey not data not founded: %s", err.Error())
		return e.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	zap.S().Info("Survey found")
	return e.JSON(http.StatusOK, survey)
}

func (h *Handler) SaveSurveyAnswer(e echo.Context) error {
	zap.S().Info("Start saving survey answer")
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	survey := &models.Survey{}
	if err := e.Bind(survey); err != nil {
		zap.S().Errorf("Bining survey error: %s", err.Error())
		return e.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if err := h.service.Survey.SaveSurvey(ctx, survey); err != nil {
		zap.S().Errorf("Saving survey answer error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("Survey answer saved successfully")
	return e.NoContent(http.StatusOK)
}
