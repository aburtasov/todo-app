package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autorizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(autorizationHeader)
	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "empthy auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[1] == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "no token in header")
		return
	}

	// parse token

	userId, err := h.services.Autorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(userCtx, userId)

}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "user id is invalid type")
		return 0, errors.New("user id is invalid type")
	}

	return idInt, nil
}
