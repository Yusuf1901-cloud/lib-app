package handler

import (
	"net/http"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) librarianSignUp(c *gin.Context) {
	var librarian models.Librarian

	if err := c.BindJSON(&librarian); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateLibrarian(librarian)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Id: id,
	})
}

func (h *Handler) librarianSignIn(c *gin.Context) {
	var librarian signInInput

	if err := c.BindJSON(&librarian); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateLibrarianToken(librarian.Username, librarian.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		Token: token,
	})
}
