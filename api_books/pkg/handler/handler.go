package handler

import (
	"api_books/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		authors := api.Group("/authors")
		{
			authors.POST("/", h.createAuthor)          //Добавить нового автора
			authors.GET("/", h.getAllAuthors)          //Получить всех авторов
			authors.GET("/:id", h.getAuthorById)       //получить автора по его идентификатору
			authors.PUT("/:id", h.updateAuthorById)    //обновить автора по его идентификатору
			authors.DELETE("/:id", h.deleteAuthorById) //удалить автора по его идентификатору
		}

		books := api.Group("/books")
		{
			books.POST("/", h.createBook)          //Добавить новую книгу
			books.GET("/", h.getAllBooks)          //Получить все книги
			books.GET("/:id", h.getBookById)       //Получить книгу по ее идентификатору
			books.PUT("/:id", h.updateBookById)    //обновить книгу по ее идентификатору
			books.DELETE("/:id", h.deleteBookById) //Удалить книгу по ее идентификатору

			books.PUT("/:id/authors/:author_id", h.updateBookWithAuthor) //одновременно обновить сведения о книге и авторе
		}
	}

	return router
}
