package connect

import (
	"fmt"
	"log"
	"os"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func DBConnectionClient(DbConfig config.PostgresConfig) *gorm.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv(DbConfig.DbUser), os.Getenv(DbConfig.DbPassword), DbConfig.DbName)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("unable to setup db connection %s", err.Error())
	}
	return db
}
