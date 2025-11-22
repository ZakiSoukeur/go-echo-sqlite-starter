package model

import "github.com/go-starter/internal/db"

type HomeData struct {
	PageDescription string
	PageTitle       string
	Authors         []db.Author
}
