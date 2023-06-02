package handler

import (
	"log"
	"net/http"

	"todolist"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input todolist.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// call the createService method
	listId, err := h.service.CreateList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Printf("err in getUserId method: %v", err)
		return
	}
	lists, err := h.service.TodoList.GetAllLists(userId)
	if err != nil {
		log.Printf("getAllLists method err: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"allLists": lists,
	})
}

func (h *Handler) getListByID(c *gin.Context) {
}
func (h *Handler) updateList(c *gin.Context) {}
func (h *Handler) deleteList(c *gin.Context) {
}
