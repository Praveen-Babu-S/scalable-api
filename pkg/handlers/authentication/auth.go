package auth

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type AuthImplementor interface {
	SignupHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
}

type AuthServer struct {
	db *gorm.DB
}

func NewAuthImplementor(db *gorm.DB) *AuthServer {
	if db == nil {
		log.Fatalln("Invalid DB object, db is nil")
	}
	return &AuthServer{
		db: db,
	}
}
