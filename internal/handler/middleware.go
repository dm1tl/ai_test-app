package handler

import (
	"ai_test-app/internal/handler/response"
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userId              = "UserId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid authorization header")
		return
	}
	token := headerParts[1]
	id, err := h.service.Authorization.Validate(ctx, token)
	if err != nil {
		logrus.Error(err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid token")
		return
	}
	c.Set(userId, id)
}

func (h *Handler) getUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get(userId)
	if !ok {
		logrus.Error("empty UserId header")
		response.NewErrorResponse(c, http.StatusBadRequest, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int64)
	if !ok {
		logrus.Error("can not set int64 type to user's id(invalid type)")
		response.NewErrorResponse(c, http.StatusInternalServerError, "internal error")
		return 0, errors.New("can not set int64 type to user's id(invalid type)")
	}
	return idInt, nil
}
