package handler

import (
	"net/http"
	"strconv"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUsersAllBooks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userBooks, err := h.services.UserBooks.GetAllUsersBooks(userId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, userBooks)
	// end calling service
}

func (h *Handler) getUsersBookById(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.UserBooks.GetUserBookById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.BookResponse{
		Book: book,
	})
}
