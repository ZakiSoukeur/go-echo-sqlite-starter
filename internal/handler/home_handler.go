package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-starter/internal/model"
	"github.com/go-starter/internal/service"
	"github.com/labstack/echo/v4"
)

type Author struct {
	Name string `json:"name" form:"name"`
	ID   int64  `json:"id" form:"id"`
}

type Handler struct {
	authorService service.AuthorService
}

func NewHandler(authorService service.AuthorService) *Handler {
	return &Handler{
		authorService: authorService,
	}
}

func (h *Handler) GetHome(c echo.Context) error {
	authors, err := h.authorService.GetAllAuthors(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	data := model.HomeData{
		PageDescription: "This is my desc",
		PageTitle:       "title",
		Authors:         authors,
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *Handler) CreateAuthor(c echo.Context) error {
	u := new(Author)
	if err := c.Bind(u); err != nil {
		return err
	}

	author, err := h.authorService.CreateAuthor(c.Request().Context(), u.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, author)
}

func (h *Handler) GetAuthor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed to convert id")
	}

	author, err := h.authorService.GetAuthor(c.Request().Context(), int64(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, author)
}

func (h *Handler) DeleteAuthor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed to convert id")
	}

	err = h.authorService.DeleteAuthor(c.Request().Context(), int64(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, "deleted")
}
