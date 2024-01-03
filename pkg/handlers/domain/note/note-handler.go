package domain

import (
	"log"

	"github.com/jinzhu/gorm"
)

type CRUDServer struct {
	db *gorm.DB
}

func NewAuthImplementor(db *gorm.DB) *CRUDServer {
	if db == nil {
		log.Fatalln("Invalid DB object, db is nil")
	}
	return &CRUDServer{
		db: db,
	}
}
