package migrations

import (
	"github.com/Praveen-Babu-S/scalable-api/models/dbmodels"
	"github.com/jinzhu/gorm"
)

// One time migrations to arrange required schema
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&dbmodels.User{}, &dbmodels.Note{}, &dbmodels.SharedNote{})
	db.Debug().Model(&dbmodels.Note{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
