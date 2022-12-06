package config

import (
	"fmt"
	"os"

	"aice-server/entity"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConfig() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to Load Env")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(dbn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed To Connect To Database")
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Transaksi{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed To close connection")
	}
	dbSql.Close()
}
