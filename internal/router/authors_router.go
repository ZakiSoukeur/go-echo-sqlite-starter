package router

import (
	"github.com/go-starter/internal/handler"
	"github.com/go-starter/internal/service"
	"github.com/labstack/echo/v4"
)

func RegisterAuthorRoutes(e *echo.Echo, authorService service.AuthorService) {
	h := handler.NewHandler(authorService)
	e.GET("/", h.GetHome)
	e.POST("/", h.CreateAuthor)
	e.GET("/:id", h.GetAuthor)
	e.DELETE("/:id", h.DeleteAuthor)
}
