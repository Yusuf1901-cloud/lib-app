package handler

import (
	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) lendBook(c *gin.Context) {
	var giveBookRequest models.GiveBookRequest

	giveBookRequest.LibrarianId, _ = getLibrarianId(c)

	h.services.LibrarianBooks.GiveBook(&giveBookRequest)

	c.JSON(200, models.Response{
		Id: giveBookRequest.BookId,
	})
}
