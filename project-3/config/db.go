package config

import (
	"fmt"
	"log"
	"os"
	"project-3/model"
	"project-3/tool"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// local
// var (
// 	host     = "localhost"
// 	user     = "postgres"
// 	password = "postgres"
// 	dbPort   = "5432"
// 	dbname   = "hacktiv8-project-3"
// 	db       *gorm.DB
// 	err      error
// )

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbPort   = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func StartDB() {
	// using ssl
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	// without ssl
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: InitLog(),
	})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	if err = db.AutoMigrate(model.Category{}, model.Task{}); err != nil {
		log.Fatal("Error run database automigration:", err)
	}

	if err = db.AutoMigrate(model.User{}); err == nil && db.Migrator().HasTable(model.User{}) {
		var count int64
		db.Model(&model.User{}).Count(&count)

		if count == 0 {
			admin := model.User{ID: 1, Fullname: "admin", Role: "admin", Password: tool.HashPassword("adminadmin"), Email: "admin@gmail.com", CreatedAt: time.Now()}

			if err = db.Model(&model.User{}).Create(&admin).Error; err != nil {
				log.Fatal("Error create admin")
			}
		}
	} else {
		log.Fatal("Error run database automigration:", err)
	}

	fmt.Println("Database connection success.")
}

func GetDB() *gorm.DB {
	return db
}
