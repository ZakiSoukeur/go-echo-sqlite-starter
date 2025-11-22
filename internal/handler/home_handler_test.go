package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-starter/internal/db"
	"github.com/go-starter/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Small struct that matches the JSON we expect from the API.
// Adjust field names/tags if your actual JSON differs.
type apiAuthor struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// setupTestHandler initializes the real DB, queries, service and handler.
func setupTestHandler(t *testing.T) (*echo.Echo, *Handler) {
	t.Helper()

	// Initialize DB (same as in main)
	err := db.InitDB("../../example.sqlite")
	require.NoError(t, err)

	conn := db.GetDB()
	require.NotNil(t, conn)

	queries := db.New(conn)
	authorService := service.NewAuthorService(queries)

	e := echo.New()
	h := NewHandler(authorService)

	return e, h
}
func int64ToString(v int64) string {
	return fmt.Sprintf("%d", v)
}

// TestAuthorFlow: create Zaki once, then get and delete the same author.
func TestAuthorFlow(t *testing.T) {
	e, h := setupTestHandler(t)

	var createdAuthorID int64

	t.Run("create Zaki", func(t *testing.T) {
		body := `{"name":"Zaki","id":1541}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h.CreateAuthor(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, rec.Code)

		var result apiAuthor
		err = json.Unmarshal(rec.Body.Bytes(), &result)
		require.NoError(t, err)

		assert.Equal(t, "Zaki", result.Name)
		// Save the ID so we can reuse it in the next subtests
		createdAuthorID = result.ID
		require.NotZero(t, createdAuthorID)
	})

	t.Run("get Zaki by ID", func(t *testing.T) {
		require.NotZero(t, createdAuthorID, "author must be created in previous subtest")

		// Build request with path param "/:id"
		req := httptest.NewRequest(http.MethodGet, "/"+int64ToString(createdAuthorID), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(int64ToString(createdAuthorID))

		err := h.GetAuthor(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)

		var result apiAuthor
		err = json.Unmarshal(rec.Body.Bytes(), &result)
		require.NoError(t, err)

		assert.Equal(t, createdAuthorID, result.ID)
		assert.Equal(t, "Zaki", result.Name)
	})

	t.Run("delete Zaki by ID", func(t *testing.T) {
		require.NotZero(t, createdAuthorID, "author must be created in previous subtest")

		req := httptest.NewRequest(http.MethodDelete, "/"+int64ToString(createdAuthorID), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(int64ToString(createdAuthorID))

		err := h.DeleteAuthor(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "deleted")
	})
}
