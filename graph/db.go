package graph

import (
	"log"
	"os"
	"gqlgen-app/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=pes11 dbname=graphql_app port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.DBUser{}, &model.DBPost{}, &model.DBComment{})
	if err != nil {
		log.Printf("Failed to auto migrate: %v", err)
	}

	DB = db
}