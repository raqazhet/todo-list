package handler

import (
	"log"
	"net/http"
	"strconv"

	"todolist"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept json
// @Produce json
// @Param input body todolist.TodoList true "list info"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
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

// @Summary GetAll todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept json
// @Produce json
// @Success 200 {object} []todolist.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
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

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} todolist.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id [get]

func (h *Handler) getListByID(c *gin.Context) {
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
	list, err := h.service.GetListById(userId, listid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"oneList": list,
	})
}

// @Summary Update list
// @Security ApiKeyAuth
// @Tags lists
// @Description update list
// @Param input body todolist.UpdateListInput true "update handler"
// @Success 200 {error} error nil
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id  [put]

func (h *Handler) updateList(c *gin.Context) {
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
	var input todolist.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.TodoList.UpdateListInout(userId, listid, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete list
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list by list.id and user.id
// @Param id path string true "delete-list"
// @Success 200 {error} error nil
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id  [delete]

func (h *Handler) deleteList(c *gin.Context) {
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
	if err := h.service.TodoList.Delete(userId, listid); err != nil {
		logrus.Printf("error in deleteList handler")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
