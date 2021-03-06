package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"userID": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
}

func (h *Handler) getListByID(c *gin.Context) {
}

func (h *Handler) updateList(c *gin.Context) {
}

func (h *Handler) deleteList(c *gin.Context) {
}
