package connect

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func DBConnectionClient(DbConfig config.PostgresConfig, logger *slog.Logger) *gorm.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv(DbConfig.DbUser), os.Getenv(DbConfig.DbPassword), DbConfig.DbName)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		logger.Error("unable to setup db connection", err, err.Error())
	}
	return db
}
