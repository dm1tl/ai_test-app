package handler

import (
	appmodels "ai_test-app/internal/app_models"
	"ai_test-app/internal/handler/response"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input appmodels.User
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	err := h.service.Authorization.Create(ctx, input)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "couldn't create an account, try again")
		return
	}
	c.JSON(http.StatusOK, response.NewStatusResponse("you succesfully signed up!"))
}

func (h *Handler) signIn(c *gin.Context) {
	var input appmodels.SignInInput
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	token, err := h.service.Authorization.Login(ctx, input)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "couldn't sign in, try again")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
