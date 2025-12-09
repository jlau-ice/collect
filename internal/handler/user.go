package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/types/interfaces"
)

type AuthHandler struct {
	userService interfaces.UserService
}

func (h *AuthHandler) AddUser(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Print(ctx)
	c.JSON(http.StatusCreated, nil)
}
