package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//parse token
	userId, err := h.services.Authorization.ParseToken((headerParts[1]))
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
	/*
		isAdmin, err := h.services.Authorization.IsAdmin(userId)
		if err != nil {
			NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		}


		c.Set("isAdmin", isAdmin)
	*/
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	logrus.Printf("user id %d", id)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}
	return idInt, nil

}

func (h *Handler) checkAdmin(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "user not found")
		return
	}
	roleId, err := h.services.Authorization.IsAdmin(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	if roleId != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "not admin")
		return
	}

	c.Next()

}
