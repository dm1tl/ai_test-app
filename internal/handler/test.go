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

func (h *Handler) genTest(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var input appmodels.TestInput
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}
	if input.Message == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty data, please write some infornation about test")
		return
	}
	test, err := h.service.Test.Create(ctx, input)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "couldn't generate test, try again")
		return
	}
	c.JSON(http.StatusOK, test)
}

func (h *Handler) answTest(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var input appmodels.AnswersInput
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}
	err := h.service.Answer(ctx, input)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "couldn't save your answers")
		return
	}
	c.JSON(http.StatusOK, response.NewStatusResponse("your answers succesfully loaded"))
}
