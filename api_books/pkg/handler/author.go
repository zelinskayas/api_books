package handler

import (
	"api_books"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createAuthor(c *gin.Context) {
	var input api_books.Authors
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Author.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id new author": id,
	})
}

type getAllAuthorsResponse struct {
	Data []api_books.Authors `json:"data"`
}

func (h *Handler) getAllAuthors(c *gin.Context) {
	authors, err := h.services.Author.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAuthorsResponse{
		Data: authors,
	})
}

func (h *Handler) getAuthorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	author, err := h.services.Author.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *Handler) updateAuthorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	_, err = h.services.Author.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "found author: "+err.Error())
		return
	}

	var input api_books.UpdateAuthorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Author.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "update author by id ok",
	})
}

func (h *Handler) deleteAuthorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	_, err = h.services.Author.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "found author: "+err.Error())
		return
	}

	err = h.services.Author.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "delete author by id ok",
	})
}
