package handler

import (
	"net/http"
	"strconv"

	"todolist"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Printf("err in getUserId method: %v", err)
		return
	}
	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input todolist.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	itemId, err := h.service.TodoItem.CreateItem(userId, listid, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"itemId": itemId,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Printf("err in getUserId method: %v", err)
		return
	}
	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	items, err := h.service.GetAllItems(userId, listid)
	if err != nil {
		logrus.Printf("error in getAllitems handler: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"allItems": items,
	})
}

func (h *Handler) getItemByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Printf("err in getUserId method: %v", err)
		return
	}
	itemid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	item, err := h.service.TodoItem.GetItemByID(userId, itemid)
	if err != nil {
		logrus.Printf("error in handler")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"oneItem": item,
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Printf("err in getUserId method: %v", err)
		return
	}
	itemid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.service.DeleteItem(userId, itemid); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Printf("err in getUserId method: %v", err)
		return
	}
	itemid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input todolist.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateItem(userId, itemid, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
