package handler

import (
	"github.com/Yusuf1901-cloud/lib-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth_api1 := auth.Group("/librarians")
		{
			auth_api1.POST("/sign-up", h.librarianSignUp) // Done
			auth_api1.POST("/sign-in", h.librarianSignIn) // Done
		}

		auth_api2 := auth.Group("/users")
		{
			auth_api2.POST("/sign-up", h.userSignUp) // DONE
			auth_api2.POST("/sign-in", h.userSignIn) // DONE
		}
	}

	api := router.Group("/api")
	{
		librarians := api.Group("/librarians", h.librarianIdentity)
		{
			books := librarians.Group("/books")
			{
				books.POST("/", h.createBook)
				books.GET("/", h.getAllBooks)
				books.GET("/:book_id", h.getBookById)
				books.PUT("/:book_id", h.updateBook)
				books.DELETE("/:book_id", h.deleteBook)
				books.POST("/lend", h.lendBook)
			}
		}

		users := api.Group("/users", h.userIdentity)
		{
			userBooks := users.Group("/books")
			{
				userBooks.GET("/", h.getUsersAllBooks)
				userBooks.GET("/:book_id", h.getUsersBookById)
			}
		}
	}

	return router
}
