package database

import (
	"fmt"
	"log"
	"os"

	"chat-platform/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 使用 MySQL 資料庫連線字符串
	// dsn := "chat_user:userpassword@tcp(127.0.0.1:3306)/chat_platform?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf(dbUser)
		panic("Failed to connect to the database!")
	}
	fmt.Println("Database connected successfully!")
}

// CloseDB closes the database
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to close the database!")
	}
	sqlDB.Close()
	fmt.Println("Database closed successfully!")
}

// MigratedDB migrates the database
func MigratedDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Group{}, &models.GroupMembers{}, &models.GroupMessages{}, &models.PrivateMessage{}, &models.Log{}, &models.LogError{})
	if err != nil {
		panic("Failed to migrate database!")
	}
	fmt.Println("Database migrated successfully!")
}
