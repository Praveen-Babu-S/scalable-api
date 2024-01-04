package auth

import (
	"log/slog"
	"net/http"

	"github.com/jinzhu/gorm"
)

type AuthImplementor interface {
	SignupHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	AuthenticateMiddleware(next http.Handler) http.Handler
}

type AuthServer struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewAuthImplementor(db *gorm.DB, logger *slog.Logger) *AuthServer {
	if db == nil {
		logger.Error("Invalid DB object, db is nil")
	}
	return &AuthServer{
		db:     db,
		logger: logger,
	}
}
