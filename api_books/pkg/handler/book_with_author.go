package handler

import (
	"api_books"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) updateBookWithAuthor(c *gin.Context) {

	book_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid book_id param")
		return
	}

	author_id, err := strconv.Atoi(c.Param("author_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid author_id param")
		return
	}

	_, err = h.services.Book.GetById(book_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "found book: "+err.Error())
		return
	}

	_, err = h.services.Author.GetById(author_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "found author: "+err.Error())
		return
	}

	var input api_books.UpdateBooksWithAuthorsInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.BookWithAuthor.Update(book_id, author_id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "update book with author by id ok",
	})
}
