package handler

import (
	"net/http"
	"strconv"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBook(c *gin.Context) {
	librarianId, err := getLibrarianId(c)
	if err != nil {
		return
	}
	var input models.Book

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// call service

	id, err := h.services.LibrarianBooks.Create(librarianId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Id: id,
	})

}

func (h *Handler) getAllBooks(c *gin.Context) {
	librarianId, err := getLibrarianId(c)
	if err != nil {
		return
	}

	books, err := h.services.LibrarianBooks.GetAllLibrarianBooks(librarianId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.GetBooksResponse{
		Data: books,
	})
}

func (h *Handler) getBookById(c *gin.Context) {
	librarianId, err := getLibrarianId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("book_id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.LibrarianBooks.GetLibrarianBookById(librarianId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.BookResponse{
		Book: book,
	})
}

func (h *Handler) updateBook(c *gin.Context) {
	librarianId, err := getLibrarianId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateLibrarianBook(librarianId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.StatusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteBook(c *gin.Context) {
	librarianId, err := getLibrarianId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("book_id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.LibrarianBooks.DeleteLibrarianBook(librarianId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.StatusResponse{
		Status: "deleted",
	})
}
