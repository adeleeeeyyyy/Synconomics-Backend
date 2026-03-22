package config

import (
	"fmt"
	"log"
	"os"
	"time"

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

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("gagal mendapatkan instance sql.DB: ", err)
	}

	// Setup Database Connection Pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

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
