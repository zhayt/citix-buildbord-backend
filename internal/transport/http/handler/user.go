package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"innovatex-app/internal/models"
	"net/http"
)

func (h *Handler) SignUp(e echo.Context) error {
	zap.S().Info("Start sing up user")
	ctx, cancel := context.WithTimeout(e.Request().Context(), h.timeout)
	defer cancel()

	var user models.User
	if err := e.Bind(&user); err != nil {
		zap.S().Errorf("Binding user data error: %s", err.Error())
		return e.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if err := h.service.User.CreateUser(ctx, &user); err != nil {
		zap.S().Error("Creating user error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("User saved successfully")
	return e.JSON(http.StatusOK, user)
}

func (h *Handler) SignIn(e echo.Context) error {
	zap.S().Info("Start sign in")
	ctx, cancel := context.WithTimeout(e.Request().Context(), h.timeout)
	defer cancel()

	var user models.User
	if err := e.Bind(&user); err != nil {
		zap.S().Errorf("Binding user data error: %s", err.Error())
		return e.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	userFromData, err := h.service.User.GetUser(ctx, &user)
	if err != nil {
		zap.S().Errorf("Getting user error: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	zap.S().Info("User entered successfully")
	return e.JSON(http.StatusOK, userFromData)
}
