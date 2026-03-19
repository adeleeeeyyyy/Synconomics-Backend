package config

import (
	"fmt"
	"log"
	"os"

	"Synconomics/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("gagal connect database: ", err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.AIMessage{},
		&models.AIResult{},
		&models.AISession{},
		&models.Business{},
		&models.Expense{},
		&models.Product{},
		&models.SupplyMatch{},
		&models.SupplyOffer{},
		&models.SupplyRequest{},
		&models.Transaction{},
		&models.TransactionItem{},
        &models.Thread{},
        &models.Reply{},
        &models.ProductSearchLog{},
        &models.MarketTrend{},
		&models.BusinessMetric{},
		&models.BusinessScore{},
	)

	DB = db
	log.Println("database terhubung")
}
