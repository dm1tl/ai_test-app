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
		response.NewErrorResponse(c, http.StatusBadRequest, "empty data, please write some information about test")
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "couldn't identificate user")
		return
	}
	test, err := h.service.Test.Create(ctx, userId, input)
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
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "couldn't identificate user")
		return
	}
	var input appmodels.AnswersInput
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}
	id, err := h.service.Answer(ctx, userId, input)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "couldn't save your answers")
		return
	}
	c.JSON(http.StatusOK, map[string]int64{
		"test_id": id,
	})
}
