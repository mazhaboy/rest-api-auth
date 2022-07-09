package handler

import (
	"github.com/gin-gonic/gin"
	rest_api_auth "github.com/mazhaboy/rest-api-auth"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input rest_api_auth.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}