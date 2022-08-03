package db

import (
	"fmt"
	"log"

	"github.com/taiki-nd/scout_go_api/config"
	"github.com/taiki-nd/scout_go_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
 * ConnectToDb
 * dbに接続するための設定
 */
func ConnectToDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		config.Config.User, config.Config.Password,
		config.Config.Host, config.Config.Port,
		config.Config.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	DB = db

	log.Printf("success db connection: %v", db)

	db.AutoMigrate(
		&models.Activity{},
		&models.License{},
		&models.Prefecture{},
		&models.School{},
		&models.Status{},
		&models.User{},
	)
}
