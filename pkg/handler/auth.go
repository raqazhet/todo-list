package handler

import (
	"log"
	"net/http"

	"todolist"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body todolist.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input todolist.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		log.Printf("error in the %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Sign-in
// @Tags auth
// @Description login in account
// @ID login
// @Accept json
// @Produce json
// @Param input body signIn true "login account"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]

func (h *Handler) signIn(c *gin.Context) {
	var input signIn
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		log.Printf("error in the %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
