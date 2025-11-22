package handler

import (
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

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	authorService service.AuthorService
}

func NewHandler(authorService service.AuthorService) *Handler {
	return &Handler{
		authorService: authorService,
	}
}

// GetHome godoc
// @Summary      Get home data
// @Description  Returns page metadata and all authors
// @Tags         authors
// @Produce      json
// @Success      201  {object}  model.HomeData
// @Failure      500  {object}  ErrorResponse
// @Router       / [get]
func (h *Handler) GetHome(c echo.Context) error {
	authors, err := h.authorService.GetAllAuthors(c.Request().Context())
	if err != nil {
		// runtime shape might differ, this is fine for docs
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	data := model.HomeData{
		PageDescription: "This is my desc",
		PageTitle:       "title",
		Authors:         authors,
	}

	return c.JSON(http.StatusCreated, data)
}

// CreateAuthor godoc
// @Summary      Create author
// @Description  Create a new author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        author  body      Author         true  "Author payload"
// @Success      201     {object}  Author
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       / [post]
func (h *Handler) CreateAuthor(c echo.Context) error {
	u := new(Author)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request body"})
	}

	author, err := h.authorService.CreateAuthor(c.Request().Context(), u.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, author)
}

// GetAuthor godoc
// @Summary      Get author by ID
// @Description  Returns a single author by ID
// @Tags         authors
// @Produce      json
// @Param        id   path      int   true  "Author ID"
// @Success      200  {object}  Author
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /{id} [get]
func (h *Handler) GetAuthor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "failed to convert id"})
	}

	author, err := h.authorService.GetAuthor(c.Request().Context(), int64(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "user not found"})
	}

	return c.JSON(http.StatusOK, author)
}

// DeleteAuthor godoc
// @Summary      Delete author by ID
// @Description  Deletes a single author by ID
// @Tags         authors
// @Produce      json
// @Param        id   path      int   true  "Author ID"
// @Success      200  {string}  string  "deleted"
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /{id} [delete]
func (h *Handler) DeleteAuthor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "failed to convert id"})
	}

	err = h.authorService.DeleteAuthor(c.Request().Context(), int64(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "user not found"})
	}

	return c.JSON(http.StatusOK, "deleted")
}
