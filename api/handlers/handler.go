package handlers

import (
	"log"

	"github.com/fesiqp/jwtauth/api/models"
)

type Handler struct {
	DB     models.Datastorer
	Logger *log.Logger
}

func New(db *models.DB, logger *log.Logger) *Handler {
	return &Handler{
		DB:     db,
		Logger: logger,
	}
}
