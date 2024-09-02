package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Connect initializes the database connection using the configuration
func Connect() {
	LoadConfig() // Ensure the configuration is loaded

	// Access the database configuration
	dbConfig := AppConfig.Database

	// Build the connection string
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	dbConfig.User,
	// 	dbConfig.Password,
	// 	dbConfig.Host,
	// 	dbConfig.Port,
	// 	dbConfig.DBName,
	// )

	fmt.Println("DB Connection : ", dbConfig.ConnectionString)

	d, err := gorm.Open(dbConfig.Driver, dbConfig.ConnectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
